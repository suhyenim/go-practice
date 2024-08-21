# Build Stage
FROM golang:1.20-alpine AS builder

# 필요한 빌드 도구 설치
RUN apk --no-cache add build-base

# 작업 디렉토리 설정
WORKDIR /app

# 모듈 캐시를 미리 빌드하기 위해 go.mod와 go.sum 복사
COPY go.mod go.sum ./
RUN go mod download

# 전체 애플리케이션 소스 코드 복사
COPY backend .

# 애플리케이션 빌드
RUN go build -o main .

# Final Stage
FROM alpine:3.18

# 애플리케이션 실행에 필요한 의존성 설치
RUN apk --no-cache add ca-certificates

# /app 디렉토리 생성
WORKDIR /app

# 빌더 이미지에서 빌드된 바이너리 복사
COPY --from=builder /app/main .

# 애플리케이션 실행
CMD ["./main"]

# Expose the service port
EXPOSE 8080
