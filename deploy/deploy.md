## Инструкция по деплою
Настоящая инструкция показывает, как за ~15 минут поднять Telegram-бота (или другой сервис)
из Git-репозитория на «чистую» виртуалку с **Debian**, используя **GitOps**-подход (FluxCD + Kubernetes).

Проверено на Debian 12.10 (bookworm), 4vCPU/8GB RAM. Меньшие ресурсы возможны, но не тестировались.

Все действия разбиты на 3 группы и выполняются от имени root (su -):
1. Тюнинг ОС.
2. Установка инфраструктуры.
3. Настройка GitOps-деплоя.

Важно!
После шага 2.8 команда cat ~/.ssh/flux.pub выведет ключ.
Его необходимо добавить в Deploy keys вашего репозитория до того, как продолжите.
Также, на шаге 3.3 вам надо будет указать токен своего бота; сейчас там заглушка <BOT_TOKEN>.

## 1. Тюнинг ОС
Подготовка ядра Debian для работы Kubernetes: отключает swap, обновляет систему,
ставит базовые утилиты, включает модули overlay и br_netfilter,
добавляет sysctl-настройки и применяет их.
```bash
swapoff -a && \
sed -i '/swap/d' /etc/fstab && \
apt-get update && \
apt-get install -y \
  curl ca-certificates gnupg lsb-release jq openssh-server && \
modprobe overlay && \
modprobe br_netfilter && \
cat <<'EOF' >/etc/sysctl.d/99-k8s.conf
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
net.ipv4.ip_forward=1
EOF
sysctl --system
```
## 2. Установка инфраструктуры
Устанавливает **containerd**, поднимает одиночный кластер **Kubernetes**, добавляет **Calico CNI**
и контроллер **FluxCD**, подготавливая кластер к GitOps-деплою из репозитория.
```bash
# 2.1 переменные - можете указать свои предпочтения по версии k8s и подсети
export K8S_VER=1.33.0            # версия Kubernetes
export POD_CIDR=10.244.0.0/16    # подсеть Pod'ов
```
```bash
# 2.2 containerd
apt-get install -y containerd && \
mkdir -p /etc/containerd && \
containerd config default >/etc/containerd/config.toml && \
sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml && \
systemctl restart containerd
```
```bash
# 2.3 kubeadm/kubelet/kubectl
curl -fsSL https://pkgs.k8s.io/core:/stable:/v${K8S_VER%.*}/deb/Release.key | \
  gpg --dearmor -o /usr/share/keyrings/k8s.gpg && \
echo "deb [signed-by=/usr/share/keyrings/k8s.gpg] \
  https://pkgs.k8s.io/core:/stable:/v${K8S_VER%.*}/deb/ /" \
  >/etc/apt/sources.list.d/kubernetes.list && \
apt-get update && \
apt-get install -y kubelet=${K8S_VER}-1.1 kubeadm=${K8S_VER}-1.1 kubectl=${K8S_VER}-1.1 && \
apt-mark hold kubelet kubeadm kubectl
```
```bash
# 2.4 single-node control-plane
kubeadm init --cri-socket unix:///run/containerd/containerd.sock \
  --kubernetes-version=v${K8S_VER} \
  --pod-network-cidr=${POD_CIDR}

mkdir -p $HOME/.kube && \
cp /etc/kubernetes/admin.conf $HOME/.kube/config && \
chown $(id -u):$(id -g) $HOME/.kube/config
```
```bash
# 2.5 Calico CNI
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
```
```bash
# 2.6 FluxCD (с toleration под single-node)
curl -s https://fluxcd.io/install.sh | bash && \
flux install --toleration-keys=node-role.kubernetes.io/control-plane
```
```bash
# 2.7 SSH-ключ для Flux (GitHub read-only)
ssh-keygen -t ed25519 -f ~/.ssh/flux -N '' && \
ssh-keyscan github.com >> ~/.ssh/known_hosts && \
kubectl create secret generic flux-git-readonly -n flux-system \
  --from-file=identity=$HOME/.ssh/flux \
  --from-file=identity.pub=$HOME/.ssh/flux.pub \
  --from-file=known_hosts=$HOME/.ssh/known_hosts
```
```bash
# 2.8 Скопируйте вывод (ключ) и добавьте в GitHub▸[ваш_проект/форк]▸Settings▸Deploy keys▸«Read-only»
cat ~/.ssh/flux.pub
```
```bash
# 2.9 Un-taint control-plane node (single-node)
# Выполняйте, если хотите запускать обычные Pods на той же VM.
kubectl taint nodes $(hostname) node-role.kubernetes.io/control-plane-
```
## 3. Настройка GitOps-деплоя
Создаёт изолированную среду (`namespace`), секрет с бот-токеном и настраивает Flux для отслеживания репозитория.
```bash
# 3.1 Переменные контекста  
# Меняйте только `APP` и `ENV`, — остальные вычисляются сами!
export APP=exrubbot ENV=prod       # prod, test, stage …
export NS=${APP}-${ENV}            # exrubbot-prod
export GIT_URL=ssh://git@github.com/akelbikhanov/${APP}.git
export KUS_PATH=./deploy/${NS}
```
```bash
# 3.2 namespace + Git-источник
kubectl create ns $NS && \
flux create source git $APP \
  --url=$GIT_URL \
  --branch=main \
  --interval=1m \
  --secret-ref=flux-git-readonly
```
```bash
# 3.3 секрет с токеном бота
kubectl create secret generic ${APP}-secrets \
  -n $NS \
  --from-literal=EXRUBBOT_TOKEN='<BOT_TOKEN>'
```
```bash
# 3.4 Kustomization (правило развертывания)
flux create kustomization $APP \
  --source=GitRepository/$APP \
  --target-namespace=$NS \
  --path=$KUS_PATH \
  --prune=true \
  --interval=10m
```