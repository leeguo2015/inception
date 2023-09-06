<template >
    <div class="container">
        <div class="centered-div">
            <div class="form_content">
                <el-form ref="form" :model="form" label-width="50px">
                <el-form-item class="input-item"  label="姓名">
                    <el-input class="input-size el-input--round"   :style="input_stayle"  v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item class="input-item"  label="密码">
                    <el-input class="input-size"  v-model="form.password"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">登录</el-button>
                    <!-- <el-button>注册</el-button> -->
                </el-form-item>
            </el-form>
    
            </div>
        </div>
    </div>
</template>

<script >

export default {
    data() {
        return {
            msg: 'ADAAAAAAAAAA ',
            input_stayle: "",
            form: {
                name: '',
                password: '',
                date1: '',
                date2: '',
                delivery: false,
                type: [],
                resource: '',
                desc: ''
            }
        }
    },
    methods: {
      onSubmit() {
        console.log('submit!');
      }
    }
}

</script>

<style>
.input-item {
    margin: 30px 0px;
}
.input-size{
    width: 10px;
    height: 40px;
    border-radius: 20px;
}

.el-form-item__label{
    color:#f56c6c;
    font-weight:bold;
    font-size:1rem !important;
    
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
    font-size:15rem;
}
</style>