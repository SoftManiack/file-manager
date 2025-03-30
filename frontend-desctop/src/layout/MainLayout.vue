<script lang="ts" setup>
    import { RouterView } from 'vue-router';
    import Typography from '../components/shared/Typography.vue'
    import Icon from '../components/shared/Icon.vue'    
</script>   

<template>
    <div class="maim-layout">
        <nav class="nav">
            <div class="nav__container px-3">
                <Typography 
                    :tagName="'h5'"
                    :color="'base'"
                > 
                    Вход
                </Typography>
                <Icon 
                    :type="'search'"
                />
                <Icon 
                    :type="'cloud'"
                />
                <Icon 
                    :type="'open-folder-silver'"
                />
                <Icon 
                    :type="'folder-silver'"
                />
            </div>
            <footer class="footer px-2">
                <Typography  :tagName="'h5'"> 
                    Объем
                </Typography>
                <div class="size-bar"></div>
            </footer>
        </nav>
        <main>
            <RouterView /> 
        </main>
    </div>
</template>     

<style lang="scss">
    @use "../assets/styles/breakpoint.scss" as *;

    .maim-layout{
        display: flex;
        height: 100%;   
        .nav{
            display: flex;
            flex-direction: column;
            max-width: 30%;
            background-color: var(--nav-bg-color);
            &__container{ 
                height: 100%;
            }
            .footer{
                border-top: 2px solid #fafafa;
                height: 7vh;
            }
        }
        main{
            height: 100% !important;
            width: 100%;
        }
    }

    @media screen and (min-width: $breakpoint-xs) and (max-width: $breakpoint-s) {
      .nav {
        width: 12%; /* На экранах от 414px до 576px */
      }
    }

    @media screen and (min-width: $breakpoint-s) and (max-width: $breakpoint-m) {
      .nav {
        width: 12%; /* На экранах от 576px до 768px */
      }
    }

    @media screen and (min-width: $breakpoint-m) and (max-width: $breakpoint-l) {
      .nav {
        width: 20%; /* На экранах от 768px до 992px */
      }
    }

    @media screen and (min-width: $breakpoint-l) and (max-width: $breakpoint-xl) {
      .nav {
        width: 20%; /* На экранах от 992px до 1200px */
      }
    }

    @media screen and (min-width: $breakpoint-xl) {
      .nav {
        width: 15%; /* На экранах больше или равных 1200px */
      }
    }
</style>