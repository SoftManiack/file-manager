<script lang="ts" setup>

    import { ref, computed, reactive  }  from 'vue'
    import { storeToRefs } from 'pinia'
    import { useVuelidate } from '@vuelidate/core'
    import { required } from '@vuelidate/validators'

    import { useAuthStore } from '../store/auth.sore'
    
    import Input from '../components/shared/Input.vue'
    import Button from '../components/shared/Button.vue'
    import Typography from '../components/shared/Typography.vue'
    
    const auth = useAuthStore() 
    const { getAuthError } = storeToRefs(auth)

    const state = reactive({
        email: null,
        password: null
    })

    const email = ref('')
    const password = ref('')
     
    const validation = computed(() => ({
        email: {
            required,
        },
        password: {
            required,
        }
    }))

    const v = useVuelidate(validation, state)
    
    const eventForButton = () => {

        auth.login( {
            email: state.email.value,
            password: state.password.value
        })  
    }
    
</script>

<template>

    <form class="login-from px-4">
        <Typography :tagName="'h4'"> 
            Вход
        </Typography>
        {{v.email.$invalid}}
        <Input 
            class="mt-4"
            v-model:modelValue="state.email"
            label="Логин"
            name="name"
            placeholder="Введите логин"
            :error="v.email.$invalid && 'Поле недолжно быть пустым'"

        />
        <Input 
            class="mt-4"
            v-model:modelValue="state.password"
            label="Пароль"
            name="password"
            placeholder="Введите пароль"
            :inputType="{ type: 'password' }"
            :error="v.password.$invalid && 'Поле недолжно быть пустым'"
        />
        <div class="login-from__error">
            {{ getAuthError.error }}
        </div>  
        <Button 
            class="mt-4"
            label="Войти"
            icon="search"
            color="green"
            type="button"
            size="m"
            @click="eventForButton"
            :error="v.email.$invalid"
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
            margin-top: 2rem !important;
        }
    }
</style>