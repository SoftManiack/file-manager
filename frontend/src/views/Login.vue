<template>
  <b-container fluid>
    <b-row class="l-auth" align-h="center">
      <b-col xl="3" lg="4" md="5" sm="7" xs="8" class="c-form">
        <div class="c-form_messange" :class="{'c-form_messange--error': Error, signInError}">
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
          type="password"
          class="mt-3" 
          v-model="form.password"
        />
        <button class="mt-3 mb-4" @click="Login" >Войти</button>
        <hr>
        <a href="/registration" >Зарегистрировать аккаунт</a>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
    name: "Login",
    data(){
      return {
        form: {
          email: "",
          password: "",
        },
        error: "",
      }
    },
    computed: {
      ...mapGetters({
        signInError: 'errorAuth',
        uidUser: "uidUser",
      }),
      Error(){
        return this.error
      }
    },
    methods: {
      ...mapActions({
        signIn: "signIn",
      }),
      async Login(){
        if (!this.form.password || this.form.login){
          this.form.password ? this.error = "Отсутствует адрес электронной почты" : this.error ="Введите пароль";
        }else{
          await this.signIn(this.form)
          this.error = this.signInError
          
          if ( !this.signInError ){
            this.$router.push(`/fm/${this.uidUser}`);
          } 
        }
      }
    },
}
</script>

<style>
</style>
