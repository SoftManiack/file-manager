<script lang="ts" setup>

    import { ref, computed }  from 'vue'
    import { storeToRefs } from 'pinia'
    import { useVuelidate } from '@vuelidate/core'
    import { required } from '@vuelidate/validators'

    import { useAuthStore } from '../store/auth.sore'
    
    import Input from '../components/shared/Input.vue'
    import Button from '../components/shared/Button.vue'
    import Typography from '../components/shared/Typography.vue'
    
    const auth = useAuthStore() 
    const { getAuthError } = storeToRefs(auth)

    const email = ref('')
    const password = ref('')
    
    const validation = computed(() => ({
        login: {
            required: required
        }
    }))
    
    const v = useVuelidate(validation, { email })
    
    const eventForButton = () => {

        auth.login( {
            email: email.value,
            password: password.value
        })
    }
</script>

<template>

    <form class="login-from px-4">
        <Typography :tagName="'h4'"> 
            Вход
        </Typography>
        <Input 
            class="mt-4"
            v-model:modelValue="email"
            label="Логин"
            name="name"
            placeholder="Введите логин"
   
        />
        <Input 
            class="mt-4"
            v-model:modelValue="password"
            label="Пароль"
            name="password"
            placeholder="Введите пароль"
        />
        <span class="login-from__error">
            {{ getAuthError }}
        </span>
        <Button 
            class="mt-4"
            label="Войти"
            icon="search"
            color="green"
            type="button"
            size="m"
            @click="eventForButton"
        />
    </form>
</template>

<style lang="scss">
    .login-from{
        margin: auto;
        width: 500px;
        margin-top: 16rem;
        background-color: var(--bg-color-form);
        box-shadow: var(--shadow-default-xs);
        border-radius: 4px;
        padding-bottom: 3rem;
        &__error{

        }
    }
</style>