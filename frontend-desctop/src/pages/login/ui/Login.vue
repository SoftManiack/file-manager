<script lang="ts" setup>

    import { ref, computed }  from 'vue'
    import { storeToRefs } from 'pinia'
    import { useVuelidate } from '@vuelidate/core'
    import { required } from '@vuelidate/validators'

    import { useAuthStore } from '@/entities/Auth/model/stores'
    
    import { Input } from '@/shared/ui/input'
    import { Button } from '@/shared/ui/button'
    import { Typography } from '@/shared/ui/typography'
    
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
        margin-top: 20rem;
        background-color: var(--bg-color-form);
        box-shadow: var(--shadow-default-xs);
        border-radius: 4px;
        &__error{

        }
    }
</style>