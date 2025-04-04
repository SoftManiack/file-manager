<script setup lang="ts">

    import { useSlots } from 'vue'
    
    interface InputType { type: 'text' | 'password' | 'number' | 'email' }

    interface IInput {
        modelValue?: string | number
        label?: string
        name?: string
        isDisabled?: boolean
        placeholder?: string
        inputType?: InputType
        error?: string
    }
    
    const props = defineProps<IInput>();
    const slots = useSlots();
    
    const emit = defineEmits(['update:modelValue'])

    const updateValue = (e) => {
        emit('update:modelValue', e.target.value)
    }
</script>

<template>
    <div class="input">
        <label
            v-if="label"
            class="input__label label"
            :for="name"
        >
            {{ label }}
        </label>
        <div class="input__container">
            <span v-if="slots.leftIcon" class="icon">
                <slot name="leftIcon"></slot>
            </span>
            <input
               :class="{'input--is-left-slot': slots.leftIcon, 'input--error': error }"
                ref="input"
                :value="modelValue"
                :type="inputType"
                :placeholder="placeholder"
                @input="updateValue"
            >
            <span v-if="slots.rightIcon" class="btn__icon">
                <slot name="rightIcon"></slot>
            </span>
        </div>
        <div
            class="input__error mt-1"
            v-if="error"
        >
            {{ error }}
        </div>
    </div>
</template>

<style lang="scss">
    .input{
        &__label{

        }
        &__container{
            position: relative;
            display: flex;
            align-items: center;
            .icon{
                position: absolute;
                margin-left: 8px;
                margin-top: 5px;
            }
            input{
                margin-top: 4px !important;
                width: 100%;
                font-size: 1rem;
                padding: 0.7rem;
                padding-left: 1rem;
                border: 1px solid var(--br-color-input);
                border-radius: 4px;
                background-color: var(--bg-color-input);
                
            }
            input:focus{
                border-color: var(--br-color-input-focus)
            }
            input:hover{
                transition: 1s all;
                border-color: var(--br-color-input-focus);
            }
            .input--is-left-slot{
                padding-left: 1.8rem;   
            }
            .input--error{
                border-color: var(--br-color-error-input) !important;
                color: var(--br-color-error-input) !important
            }
            .input--error:hover{
                border-color: var(--br-color-error-input) !important;
                color: var(--br-color-error-input) !important
            }
        }
        &__error{
            color: var(--text-color-error-input) !important;
            font-size: 0.9rem;
        }
    }
</style>