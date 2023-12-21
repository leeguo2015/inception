<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-22 00:08:59
 * @FilePath: \inception\web\src\views\LoginView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template >
    <div class="container">
        <div class="centered-div">
            <div class="form_content">
                <el-form ref="form" :model="form" label-width="50px">
                    <el-form-item class="input-item" label="姓名">
                        <el-input class="input-size el-input--round" :style="input_stayle" v-model="form.name"></el-input>
                    </el-form-item>
                    <el-form-item class="input-item" label="密码">
                        <el-input class="input-size" v-model="form.password"></el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">登录</el-button>
                        <!-- <el-button>注册</el-button> -->
              
                    </el-form-item>
                </el-form>

            </div>
        </div>
        <button @click="showMessageBox">显示提醒框</button>
    </div>
</template>
<script>
import { ElMessageBox } from 'element-plus';

import { ElMessage } from 'element-plus';


export default {
    data() {
        return {
            msg: '',
            input_stayle: "",
            form: {
                name: 'leeguo',
                password: '123456',
            }
        }
    },
    methods: {
        onSubmit() {
            const formData = new FormData();
            formData.append('username', this.form.name);
            formData.append('password', this.form.password);
            this.$api.post("/user/login", formData
            ).then(res => {
                if (res.code != "200") {
                    this.showMessageBox(res.msg)
                }
            }).catch(err => {
                
            })
        },
        showMessageBox(msg) {
      ElMessageBox.alert(msg, '登录失败', {
        confirmButtonText: '确定',
        type: 'error'
      });
    }


    }
}

</script>

<style>
.input-item {
    margin: 30px 0px;
}

.input-size {
    width: 10px;
    height: 40px;
    border-radius: 20px;
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
    padding: 50px 50px;
    width: 500px;
    height: 600px;
    background-color: rgba(255, 255, 255, 1);
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    font-size: 15rem;
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
</style>