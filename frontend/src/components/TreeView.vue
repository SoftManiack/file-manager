<template>
    <div>
        <div v-for="item in data" v-bind:key="item.uid">
            <router-link class="c-nav_item pl-2 pt-1" :to="{ name: 'Main', params: { uid: uidUser }}" @click.native="toggleActive('disk')">
                <div v-if="item.disclosed" @click="toggleClose">
                    <svg class="ml-1 mr-1 c-nav_arrow" height="8" width="8"><use xlink:href="@/assets/icons/sprite.svg#right-arrow"></use></svg>
                </div>
                <div v-if="!item.disclosed" @click="toggleClose">
                    <svg class="ml-1 mr-1 c-nav_arrow" height="8" width="8"><use xlink:href="@/assets/icons/sprite.svg#bottom-arrow"></use></svg>
                </div>
                <svg class="c-nav_icon mr-2" width="16" height="14" :fill ="activeDisk"><use xlink:href="@/assets/icons/sprite.svg#cloud"></use></svg>
                {{ item.name }}
            </router-link>
            <div v-if="item.disclosed">
                <TreeViewChild
                    :child="item.child"
                    :lvl="1"
                />
            </div>
            <router-link class="c-nav_item px-4" :to="{ name: 'Recent' }" @click.native="toggleActive('recent')">
                <svg class="c-nav_icon mr-2" width="17" height="15" :fill ="activeRecent" ><use xlink:href="@/assets/icons/sprite.svg#time"></use></svg>
                Недавние
            </router-link>
            <router-link class="c-nav_item px-4" :to="{ name: 'Trash' }" @click.native="toggleActive('trash')">
                <svg class="c-nav_icon mr-2" width="17" height="16" :fill ="activeTrash" ><use xlink:href="@/assets/icons/sprite.svg#trash"></use></svg>
                Корзина
            </router-link>
        </div>
    </div>
</template>

<script>
import TreeViewChild from '@/components/TreeViewChild.vue'
import { mapGetters } from "vuex";

export default {
    data() {
        return {
            active: {
                disk: true,
                recent: false,
                trash: false,
            },
            data: [
                {
                    name: 'Облачный диск',
                    uid: '12ssf2bfd1',
                    disclosed: false,
                    child: [
                        {
                            name: 'Папка11',
                            uid: '121sfs1sd',
                            disclosed: false,
                            child: []
                        },
                        {
                            name: 'Папка12',
                            uid: '12sfs1sd',
                            disclosed: false,
                            child: [
                                {
                                    name: 'Папка121',
                                    uid: '12dsfs1sd',
                                    disclosed: false,
                                    child: []
                                },
                                {
                                    name: 'Папка122',
                                    uid: '121sfs1sd',
                                    disclosed: false,
                                    child: [
                                        {
                                            name: 'Папка121',
                                            uid: '12dsfs1sd',
                                            disclosed: false,
                                            child: []
                                        },
                                    ]
                                }
                            ]
                        },
                        {
                            name: 'Па3',
                            uid: '13s1sd',
                            disclosed: false,
                            child: [
                                {
                                    name: 'П12ка121',
                                    uid: '12sxsfs1sd',
                                    disclosed: false,
                                    child: []
                                },
                                {
                                    name: 'П1пка122',
                                    uid: '112sfs1sd',
                                    disclosed: false,
                                    child: []
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    },
    components: {
        TreeViewChild,
    },
    computed: {
        ...mapGetters({
            uidUser: "uidUser",
        }),
        activeDisk(){
            return this.active.disk ? '#2ba6de': '#6b6b6b';
        },
        activeRecent(){
            return this.active.recent ? '#2ba6de': '#6b6b6b';
        },
        activeTrash(){
            return this.active.trash ? '#2ba6de': '#6b6b6b';
        }
    },
    methods: {
        toggleClose(){
            this.data[0].disclosed = !this.data[0].disclosed;
        },
        toggleActive(active){
            console.log(active)
            if(active == 'disk'){
                this.active.disk = true;
                this.active.recent = this.active.trash = false;
            }
            if(active == 'recent'){
                this.active.recent = true;
                this.active.disk = this.active.trash = false;
            }
            if(active == 'trash'){
                this.active.trash = true;
                this.active.recent = this.active.disk = false;
            }
        }
    }
}
</script>

<style>
</style>




