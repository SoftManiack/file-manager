<template>
  <b-container fluid>
    <b-row class="l-auth" align-h="center">
      <b-col xl="3" lg="4" md="5" sm="7" xs="8" class="c-form">
        <div class="c-form_messange" :class="{'c-form_messange--error': Error, signUpError}">
          <p>{{ error }}</p>
        </div>
        <h6>Вход в FileManager</h6>
        <input
          placeholder="Укажите адрес электронной почты" 
          class="mt-4"
          v-model="form.email"
        />
        <input 
          placeholder="Введите пароль" 
          class="mt-3" 
          v-model="form.password"
        />
         <input 
          placeholder="Повторите пароль" 
          class="mt-3" 
          v-model="form.passwordRepeat"
        />
        <button class="mt-3 mb-4" @click="Registration" >Регистрация</button>
        <hr>
        <a href="/login" >Войти в FileManager</a>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
    name: "Registration",
    data(){
      return {
        form: {
          email: "",
          password: "",
          passwordRepeat: "",
        },
        error: "",
      }
    },
    computed: {
      ...mapGetters({
        signUpError: 'errorAuth',
      }),
      Error(){
        return this.error
      }
    },
    methods: {
      ...mapActions({
        signUp: "signUp",
      }),
      async Registration(){
        if (!this.form.password || !this.form.email){
          this.form.password ? this.error = "Отсутствует адрес электронной почты" : this.error ="Введите пароль";
        }
        if (this.form.password !== this.form.passwordRepeat){
          this.error = "Пароли должны совпадать";
        }
        if (this.form.password === this.form.passwordRepeat && this.form.password && this.form.email){
          await this.signUp(this.form)
          this.error = this.signUpError
          if ( !this.signUpError ){
             this.$router.push('/login');
          } 
        }
      }
    },
}
</script>

<style>
</style>
