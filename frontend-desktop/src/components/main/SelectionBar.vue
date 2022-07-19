<template>
    <div class="selection-bar">
        {{count}}
    </div>
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