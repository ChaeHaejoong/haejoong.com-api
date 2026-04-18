## 인프라
- [x] Docker + Air + Log Volume 세팅
- [ ] PostgreSQL 연동 및 Gorm 설정
- [ ] Redis 컨테이너 추가 및 Go 클라이언트(go-redis) 연동
- [ ] Redis 기반 캐싱/세션 전략 수립 (TTL 정책 등)

## Phase 1: Auth & Middleware
- [ ] JWT 인증 미들웨어 구현 (Access/Refresh)
- [ ] Refresh Token 쿠키 정책 구현 (httpOnly, secure, sameSite, maxAge)
- [ ] 공통 에러 핸들러 및 JSON 응답 구조 설계
- [ ] 인증/인가 미들웨어 체인 설계 (Optional Auth 포함)
- [ ] 구조화된 로깅(Structured Logging) 적용 (zap 또는 slog 활용)
- [ ] 요청 단위 로깅/트레이싱 미들웨어 추가

### Endpoints (Auth)
- [ ] `POST /auth/register` - 회원가입
- [ ] `POST /auth/login` - 로그인 + refreshToken 쿠키 발급
- [ ] `POST /auth/logout` - 로그아웃 + refreshToken 제거
- [ ] `POST /auth/refresh` - accessToken 재발급 + refreshToken 로테이션
- [ ] `POST /auth/verify-password` - 비밀번호 검증 (인증 필요)

### 추가로 해야할 점
- [ ] JWT 키/만료시간 환경변수 전략 확정
- [ ] Redis 기반 Refresh Token 저장 및 로테이션 정책 확정
- [ ] Cookie SameSite 정책(개발/운영) 분리 설계
- [ ] 에러 코드 표준화 (`code`, `message`, `details`, `requestId`)


## Phase 2: Users & Tags
- [ ] User 도메인 모델/리포지토리 마이그레이션
- [ ] Role 기반 인가(Admin/User) 구현
- [ ] 페이지네이션 공통 유틸 구현
- [ ] soft delete 정책 반영

### Endpoints (Users & Tags)
- [ ] `GET /users/check-userid` - 아이디 중복 확인
- [ ] `GET /users/check-nickname` - 닉네임 중복 확인
- [ ] `GET /users/me` - 내 정보 조회 (인증 필요)
- [ ] `GET /users` - 유저 목록 (Admin)
- [ ] `GET /tags` - 태그 목록
- [ ] `POST /tags` - 태그 생성 (Admin)

### 추가로 해야할 점
- [ ] soft delete 시 unique 컬럼(아이디/닉네임) 처리 정책 확정
- [ ] 페이지네이션 응답 포맷 표준화 (`items`, `total`, `page`, `pageSize`)


## Phase 3: Posts & Comments
- [ ] Post/Comment 도메인 구현 및 쿼리 최적화
- [ ] 공개/비공개 접근 정책 및 좋아요/조회수 처리
- [ ] Redis 기반 조회수 중복 방지 (IP/User별 쿨타임)

### Endpoints
- [ ] `GET /posts/published` - 공개 글 목록 (Optional Auth)
- [ ] `GET /posts/:id` - 글 상세 (Optional Auth)
- [ ] `POST /comments/posts/:postId` - 댓글 생성 (인증 필요)
- [ ] `POST /posts/:id/like` - 좋아요 토글 (인증 필요)

### 추가로 해야할 점
- [ ] 조회수 버퍼링 전략 (Redis 카운트 -> DB 배치 쓰기) 검토
- [ ] Optional Auth 실패 시 graceful fallback 정책 정의


## Phase 4: Images & Storage (Garbage Collection)
- [ ] 파일 업로드 파이프라인 (multipart/form-data)
- [ ] 로컬/S3 스토리지 추상화 마이그레이션
- [ ] 이미지 유예 삭제 로직 구현
    - [ ] `linked_at` 컬럼이 Null인 미사용 이미지 식별
    - [ ] 7일간 연결되지 않은 이미지/DB 레코드 자동 삭제 워커

### Endpoints
- [ ] `GET /images/mine` - 내 이미지 목록 (인증 필요)
- [ ] `POST /images/upload` - 이미지 업로드 (인증 필요)

### 추가로 해야할 점
- [ ] 이미지 정리 정책: 게시글 저장 실패로 고아가 된 이미지(Orphaned Images) 처리 로직 확정
- [ ] 삭제 워커 주기 설정 (매일 새벽 Cron 작업 등)


## Phase 5: Quality & Security
- [ ] 통합 테스트(E2E) 작성
- [ ] OpenAPI(Swagger) 문서 자동화
- [ ] Rate limit/CORS/보안 헤더 적용 (Redis 활용)
- [ ] 헬스체크/메트릭 연동

---

### Migration Note
- **Redis 활용**: Refresh Token, Rate Limit, 조회수 중복 방지, 미사용 이미지 삭제 트리거
- **Worker**: 주기적인 DB 스캔을 통해 7일 이상 방치된 `unlinked` 이미지 자동 정리