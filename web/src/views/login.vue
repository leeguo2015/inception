<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-23 01:38:14
 * @FilePath: \inception\web\src\views\LoginView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="container">
    <div class="centered-div">
      <el-card class="login-card">
        <div class="card-top">
          <div class="logo-wrap">
            <div class="logo-icon">IC</div>
            <div class="logo-text">Inception</div>
          </div>
          <div class="card-desc">内容创作 · 简洁高效的博客平台</div>
        </div>

        <el-tabs v-model="activeTab" type="card">
          <el-tab-pane label="登录" name="login">
            <el-form ref="form" :model="form" label-width="0px">
              <el-form-item class="input-item">
                <el-input v-model="form.name" class="input-size" placeholder="用户名">
                  <template #prefix>
                    <i class="el-icon-user"></i>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item class="input-item">
                <el-input v-model="form.password" class="input-size" type="password" placeholder="密码">
                  <template #prefix>
                    <i class="el-icon-lock"></i>
                  </template>
                </el-input>
              </el-form-item>

              <div class="actions-row">
                <el-checkbox v-model="remember">记住我</el-checkbox>
                <a class="forgot" @click.prevent="forgot">忘记密码?</a>
              </div>

              <el-form-item>
                <el-button type="primary" @click="onSubmit" class="primary-btn">登录</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>

          <el-tab-pane label="注册" name="register">
            <el-form ref="regForm" :model="registerForm" label-width="0px">
              <el-form-item class="input-item">
                <el-input v-model="registerForm.name" class="input-size" placeholder="用户名">
                  <template #prefix>
                    <i class="el-icon-user"></i>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item class="input-item">
                <el-input v-model="registerForm.email" class="input-size" placeholder="邮箱">
                  <template #prefix>
                    <i class="el-icon-message"></i>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item class="input-item">
                <el-input v-model="registerForm.password" class="input-size" type="password" placeholder="密码">
                  <template #prefix>
                    <i class="el-icon-lock"></i>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="onRegister" class="primary-btn">创建账户</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>

        <div class="card-footer">
          <div class="foot-text">或使用第三方账号登录</div>
          <div class="social-icons"> 
            <i class="el-icon-s-platform"></i>
            <i class="el-icon-s-promotion"></i>
          </div>
        </div>
      </el-card>
    </div>
  </div>
</template>
<script>
import { ElMessageBox } from "element-plus";

import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      msg: "",
      form: {
        name: "",
        password: "",
      },
      activeTab: "login",
      registerForm: {
        name: "",
        password: "",
        email: "",
      },
      remember: false,
    };
  },
  methods: {
    onSubmit() {
      const formData = new FormData();
      formData.append("username", this.form.name);
      formData.append("password", this.form.password);
      this.$api
        .post("/user/login", formData)
        .then((res) => {
          if (res.code == "200") {
            ElMessage.success("登录成功");
            this.$store.commit("SET_TOKEN", res.data.token);
            this.$store.commit("SET_USER", res.data.userInfo);
            // 延迟0.5s跳转至主页
            setTimeout(()=>{
            this.$router.push("/blog_add");
            },800)

          } else {
            this.showMessageBox(res.msg);
          }
        })
        .catch((err) => {});
    },
    showMessageBox(msg) {
      ElMessageBox.alert(msg, "登录失败", {
        confirmButtonText: "确定",
        type: "error",
      });
    },
    onRegister() {
      const payload = {
        username: this.registerForm.name,
        password: this.registerForm.password,
        email: this.registerForm.email,
      };

      this.$api
        .post("/user/register", payload)
        .then((res) => {
          if (res.code == "200") {
            ElMessage.success("注册成功，请登录");
            this.activeTab = "login";
            this.registerForm = { name: "", password: "", email: "" };
          } else {
            ElMessageBox.alert(res.msg || "注册失败", "注册失败", { type: "error" });
          }
        })
        .catch((err) => {
          ElMessageBox.alert(err?.message || "网络错误", "注册失败", { type: "error" });
        });
    },
    forgot() {
      ElMessage.info('请联系管理员重置密码');
    },
  },
};
</script>

<style scoped>
.input-item {
  margin: 30px 0px;
}

.input-size {
  width: 100%;
  height: 40px;
  border-radius: 6px;
}

.el-form-item__label {
  color: #f56c6c;
  font-weight: bold;
  font-size: 1rem !important;
}

.login-container {
  width: 100px;
  height: 100px;
  /* margin: auto; */
  /* background-color: aqua; */
}

.centered-div {
  padding: 28px;
  width: 360px;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.08);
}

.login-card {
  border-radius: 10px;
  overflow: hidden;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 18px;
  border-bottom: 1px solid #f2f2f2;
}

.logo {
  font-weight: 700;
  font-size: 18px;
  color: #409eff;
}

.subtitle {
  font-size: 12px;
  color: #999;
}

.container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.full-btn {
  width: 100%;
}

.centered-div {
  position: relative;
  z-index: 1;
  width: 420px;
  max-width: calc(100% - 40px);
  padding: 0;
}

.login-card {
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 12px 40px rgba(16,39,112,0.18);
}

.card-top {
  padding: 22px 24px 16px 24px;
  background: linear-gradient(90deg, rgba(255,255,255,0.06), rgba(255,255,255,0.02));
}

.logo-wrap{
  display:flex;align-items:center;gap:12px;
}
.logo-icon{
  width:48px;height:48px;border-radius:10px;background:linear-gradient(135deg,#409eff,#67c23a);display:flex;align-items:center;justify-content:center;color:#fff;font-weight:700;font-size:18px;
}
.logo-text{font-size:20px;color:#fff;font-weight:700}
.card-desc{font-size:12px;color:rgba(255,255,255,0.85);margin-top:8px}

.el-tabs__header{padding:20px}
.input-item{margin:12px 0}
.input-size{width:100%;height:44px;border-radius:22px}
.actions-row{display:flex;justify-content:space-between;align-items:center;margin:8px 0 16px}
.forgot{color:rgba(0,0,0,0.45);font-size:13px;cursor:pointer}
.primary-btn{width:100%;height:46px;border-radius:24px;background:linear-gradient(90deg,#409eff,#67c23a);color:#fff;border:none}
.card-footer{padding:16px 20px;border-top:1px solid rgba(0,0,0,0.04);display:flex;justify-content:space-between;align-items:center}
.foot-text{color:rgba(0,0,0,0.45);font-size:12px}
.social-icons i{font-size:18px;color:rgba(0,0,0,0.45);margin-left:8px}
</style>
