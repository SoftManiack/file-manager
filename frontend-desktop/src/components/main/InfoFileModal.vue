<template>
    <b-modal 
        content-class="l-main_info-modal info-modal"
        id="modal-infofile"
        ref="modal-infofile"
        header-class="info-modal_header"
        hide-footer 
        centered>
        <div class="info-modal_body pt-3 pl-2">
            <b-row>
                <b-col>
                    <p class="info-modal_title">Имя: <span>{{selection[0].name + "." + selection[0].extension}}</span></p>
                </b-col>
            </b-row>
            <b-row>
                <b-col class="info-modal_path ml-2 mr-3 pt-2">
                    <p class="info-modal_title mb-0">Путь:</p>
                    <p class="mb-1">{{selection[0].path}}</p>
                </b-col>
            </b-row>
            <b-row class="mt-3">
                <b-col>
                    <p class="info-modal_title mb-0">Размер:</p>
                    <p class="mb-0">{{selection[0].size}} Б</p>
                </b-col>
                <b-col>
                    <p class="info-modal_title mb-0 ">Последнее изменение:</p>
                    <p class="mb-0">{{selection[0].date_update}}</p>
                </b-col>
            </b-row>
            <b-row>
                <b-col>
                    <p class="info-modal_title mb-0">Дата создания</p>
                    <p class="mb-1">{{selection[0].date_create}}</p>
                </b-col>
            </b-row>
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