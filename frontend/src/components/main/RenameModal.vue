<template>
    <b-modal 
        content-class="l-main_new-modal-fd new-modal-fd"
        header-class="new-modal-fd_header"
        id="modal-rename"
        ref="modal-rename"
        hide-footer 
        centered
    >
        <div class="new-modal-fd_body">
            
        </div>
    </b-modal>
</template>

<script>

export default {
    name: "RenameModal",
    emits: ['closeMenu'],
    props: {
        selection: { type: Array, default: function () {} },
    },
}
</script>

<style>
</style>