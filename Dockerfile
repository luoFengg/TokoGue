# --- STAGE 1: Build (Dapur Kotor) ---
# Kita pakai image Go yang lengkap buat compile kodingan
FROM golang:alpine AS builder

# Set folder kerja di dalam container
WORKDIR /app

# Copy file dependency dulu (biar cache-nya awet)
COPY go.mod go.sum ./

# Download semua library (gin, gorm, redis, midtrans, dll)
RUN go mod download

# Copy semua kodingan kita ke dalam container
COPY . .

# Build aplikasi jadi binary file (namanya 'main')
RUN go build -o main main.go


# --- STAGE 2: Run (Piring Saji) ---
# Kita pindahkan hasil masakan ke image Alpine yang kosong dan ringan
FROM alpine:latest

    # # Tambah di stage 2
    # RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy hasil build (binary 'main') dari Stage 1
COPY --from=builder /app/main .

# Copy file .env (Opsional: Nanti di production .env biasanya di-inject, tapi buat belajar copy aja dulu gpp)
# COPY .env . 
# (SAYA SARANKAN JANGAN COPY .ENV, KITA PASSING LEWAT DOCKER COMPOSE NANTI BIAR AMAN)

# Buka port 5000 (biar bisa diakses dari luar)
EXPOSE 5000

# Perintah buat nyalain aplikasi
CMD ["./main"]