<script setup lang="ts">

    import { useSlots } from 'vue'

    interface IButton {
        label?: string
        name?: string
        size?: 's' | 'm' | 'l' | 'xl'
        icon?: string,
        color?: string
        isDisabled?: boolean
        error?: string
    }
    
    const props = defineProps<IButton>();
    const slots = useSlots();

    const emit = defineEmits(['click'])
    const clickOnButton = () => {
        emit('click')
    }

</script>

<template>
    <button 
        :class="['btn', `btn__${size}`, `btn__${color}`]"
        @click="clickOnButton"
    > 
        <span v-if="slots.leftIcon" class="btn__icon">
            <slot name="leftIcon"></slot>
        </span>
        {{ label }} 
    </button>
</template>

<style lang="scss">
    .btn{
        cursor: pointer;
        border: 1px solid;
        padding: 0.4rem;
        border-radius: 4px;
        border: 1px solid silver;
        display: flex;
        align-items: center;
        &__icon{
          margin-right: 6px;
        }
    }
    .btn:hover{
        
    }
    .btn__green{  
        background-color: #2CBA8F;
        border-color: var(--button-green-color-border);
        color: var(--button-green-color-font);
    }
    .btn__m{
        font-size: 1rem;
        padding: 8px;
        padding-left: 15px;
        padding-right: 15px;
    }
    
</style>
