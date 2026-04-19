# NestJS to Go 마이그레이션 🐁

해중닷컴의 기존 NestJS API를 Go로 마이그레이션하기 위한 레포지토리

---

### 📌 마이그레이션 전략
NestJS의 각 요소를 Go의 특성에 맞춰 아래처럼 매핑하여 구현

| NestJS | Go(Gin) |
| :--- | :--- |
| Controller | Handler / Router |
| Service (Provider) | Services |
| DTO (Validation) | Struct with Tags |
| TypeORM Entity | Gorm Models |
| Middleware | Custom Middleware |

---

### 🖥️ 환경
- **Server**: GMKtec G6 (Ryzen 5425U / 32GB RAM / Ubuntu) **< 홈서버**
- **Runtime**: Docker Compose
- **Dev Tool**: [Air](https://github.com/air-verse/air) (Hot Reload)
- **Log Path**: `./logs/api.log`

---

### 명령어

```bash
# 1. 개발 (Hot Reload)
make dev
# 2. 배포 (백그라운드 실행)
make up
```

---

### ✅ 체크포인트
마이그레이션 진행 상황은 [PROGRESS.md](./docs/PROGRESS.md) 파일에서 실시간으로 관리