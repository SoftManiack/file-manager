<template>
    <div>
        <div v-for="item in child" v-bind:key="item.uid">
            <div v-if="item.child[0]">
                <div class="c-nav_item pl-2 pt-1" :style="paddingDir">
                    <div v-if="item.disclosed" @click="toggleClose(item.uid)">
                        <svg class="ml-1 mr-1 c-nav_arrow" height="8" width="8"><use xlink:href="@/assets/icons/sprite.svg#right-arrow"></use></svg>
                    </div>
                    <div v-if="!item.disclosed" @click="toggleClose(item.uid)">
                        <svg class="ml-1 mr-1 c-nav_arrow" height="8" width="8"><use xlink:href="@/assets/icons/sprite.svg#bottom-arrow"></use></svg>
                    </div>  
                    <div v-if="!item.disclosed">
                        <svg class="c-nav_folder" width="19" height="19"><use xlink:href="@/assets/icons/sprite.svg#folder"></use></svg>
                    </div>
                    <div v-if="item.disclosed">
                        <svg class="c-nav_folder" width="19" height="19"><use xlink:href="@/assets/icons/sprite.svg#folder-open"></use></svg>
                    </div>
                    <router-link class="ml-1 px-1" :to="{ name: 'Main', params: { id: item.uid }}">{{ item.name }} </router-link>
                </div>
            </div>
            <div v-else class="c-nav_item px-4" :style="paddingDirDisclosed">
                <svg class="c-nav_cloud" width="19" height="19"><use xlink:href="@/assets/icons/sprite.svg#folder"></use></svg>
                <router-link class="ml-2" :to="{ name: 'Main', params: { id: item.uid }}">{{ item.name }} </router-link>
            </div>
            <div v-if="item.disclosed">
                <TreeViewChild
                    :child="item.child"
                    :lvl ="LVL"
                />
            </div>
        </div>
    </div>
</template>

<script>
import TreeViewChild from '@/components/TreeViewChild.vue'

export default {
    name: "TreeViewChild",
    props: {
        child: { 
            type: Array, 
            default: () => (
                [
                    {
                        name: '',
                        uid: '',
                        disclosed: true,
                        child: []
                    }
                ]
            ) 
        },
        lvl: {
            type: Number,
            default: 1
        },
    },
    components: {
        TreeViewChild
    },
    computed: {
        LVL(){
            return this.lvl + 1;
        },
        paddingDir(){
            return `padding-left: ${ this.LVL * 11 }px !important;`;
        },
        paddingDirDisclosed(){
            return `padding-left: ${ this.LVL * 19 }px !important;`;
        }
    },
    methods: {
        toggleClose(uid){
            this.child.filter(item => item.uid == uid)[0].disclosed = !(this.child.filter(item => item.uid == uid)[0].disclosed)
        }
    }
}
</script>

<style>
</style>