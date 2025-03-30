<script lang="ts" setup>
    import { defineProps } from 'vue'

    interface Props {
        tagName?: 'p' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5',
        bold?: boolean,
        size?: 'xs' |  's' | 'm' | 'l' | 'xl',
        color?: 'base'
    }
    
    const props = defineProps<Props>();
</script>
    
<template>
    <component 
        :is="props.tagName" 
        :class="[ 'typography', size ? `${size}` : '', bold ? 'bold' : '', `tag_${tagName}`, `color-${color}`]"
    >
        <slot></slot>
    </component>
</template>

<style>
    .typography{
        font-weight: normal;
        line-height: 150%;
    }

    .xs{ font-size: 12px; }
    .s{ font-size: 16px; }
    .m{ font-size: 18px; }
    .l{ font-size: 25px; }
    .xl{ font-size: 36px; }

    .tag_h1 { font-size: 64px; }
    .tag_h2 { font-size: 48px; }
    .tag_h3 { font-size: 36px; }
    .tag_h4 { font-size: 24px; }
    .tag_h5 { font-size: 18px; }

    .color-base{ color: #333333; }

    @media screen and (max-width: 768px){
        .xs{ font-size: 8px; }
        .s{ font-size: 12px; }
        .m{ font-size: 14px; }
        .l{ font-size: 20px; }
        .xl{ font-size: 32px; } 

        .tag_h1 { font-size: 48px; }
        .tag_h2 { font-size: 36px; }
        .tag_h3 { font-size: 24px; }
        .tag_h4 { font-size: 18px; }
        .tag_h5 { font-size: 16px; }
    }

    .bold{
        font-family: "Rubik-Bold";
    }
</style>