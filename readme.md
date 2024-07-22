
<p align="center">
  <img width="192" height="192" src="https://github.com/user-attachments/assets/7eaf11c2-80f5-443a-956f-04393e7c58d5">
</p>

~~# **Goimbo** is a Go-powered textboard / imageboard engine~~

This is a fork. It's going to be kept strictly as a textboard engine. 

# Setup
It is assumed that you already have Go and Git installed, as well as the Postgres database.

1. Clone the repo
```bash
git clone https://github.com/1ort/goimbo.git
cd goimbo
```
2. Install dependencies
```bash
go mod download
```
3. Create and edit config file
```bash
cp base_config.yaml config.yaml
```
4. Build and run!
```bash
go build .
./goimbo
```

# To-Do
- [x] Read-only API
- [x] CSRF protection
- [x] Captcha
- [x] Greentexting
- [ ] Markdown file upload
- [ ] Rate-limiter
- [ ] Admin functionality
- [ ] Tripcodes
- [ ] password-based file/post deletion
