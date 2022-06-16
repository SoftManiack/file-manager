<template>
    <b-container class="selection-bar  pt-3" fluid>
        <b-row>
            <b-col>
                <span>{{ count }} элементов выбранно</span>
            </b-col>
            <b-col class="ml-auto d-flex flex-row-reverse">
                <button class="selection-bar_btn">
                    <svg width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#trash-2"></use></svg>
                </button>
                <button class="selection-bar_btn">
                    <svg width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#link"></use></svg>
                </button>
                <button class="selection-bar_btn">
                    <svg width="19" height="17"><use xlink:href="@/assets/icons/sprite.svg#donwload-2"></use></svg>
                </button>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
export default {
    name: "SelectionBar",
    props: {
        selection: { type: Array, default: function () {}},
    },
    computed: {
        count(){
            return this.selection.length
        },
        /*memory(){
            this.selection.reduce((a,b)=> a+b )
        }*/
    }
}
</script>

<style>

</style>