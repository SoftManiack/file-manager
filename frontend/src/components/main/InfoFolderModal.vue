<template>
    <b-modal 
        content-class="l-main_info-modal info-modal"
        id="modal-infofolder"
        ref="modal-infofolder"
        hide-header
        hide-footer 
        centered>
        <div class="info-modal_body">
            <div class="info-modal_left">1</div>
            <div class="info-modal_right">1</div>
        </div>
    </b-modal>
</template>

<script>
export default {
    name: "InfoFolderModal",
    emits: ['closeMenu'],
    props: {
        selection: { type: Array, default: function () {} },
    },
}
</script>

<style>
</style>