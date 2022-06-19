<template>
    <b-modal 
        content-class="l-main_new-modal-fd new-modal-fd"
        header-class="new-modal-fd_header"
        id="modal-infofile"
        ref="modal-infofile"
        hide-footer 
        centered>
        <div class="new-modal-fd_body">
            {{selection}}
        </div>
    </b-modal>
</template>

<script>

export default {
    name: "InfoFileModal",
    emits: ['closeMenu'],
    props: {
        selection: { type: Array, default: function () {} },
    },
}
</script>

<style>
</style>