<template>
    <b-modal 
      content-class="l-main_modal modal"
      header-class="modal_header"
      id="modal-rename"
      ref="modal-rename"
      hide-footer 
      centered
    >
      <div class="modal_body">
        <template v-if="selection[0].type == 'Directory'">
          <h5 class="mb-4">Переименовать папку</h5>
        </template>
        <template v-else>
          <h5 class="mb-4">Переименовать файл</h5>
        </template>
        <b-row class="px-3">
          <b-form-input 
            class="modal_input"  
            :class="{'modal_input--error': Error}"
            :placeholder="selection[0].name"
            v-model="form.newName"
            ></b-form-input>
        </b-row>
        <transition name="fade">
          <div v-if="Error" class="modal_error mt-2">
            {{ Error }}
          </div>
        </transition>
        <b-row class="mt-3 px-3" align-h="end">
          <button class="modal_exit" size="sm" @click="hideModal" >Отмена</button>
          <button class="modal_create modal_rename ml-2" size="sm" @click="rename">Переименовать</button>
        </b-row>
      </div>
    </b-modal>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
    name: "RenameModal",
    data(){
      return {
        form: { 
          newName: "",
          type: "",
          uid: "",
        },
        error: "",
      }
    },
    emits: ['closeMenu'],
    props: {
        selection: { type: Array, default: function () {} },
    },
    computed: {
      ...mapGetters({
        getElements: "elements",
      }),
      Error(){
        return this.error
      }
    },
    watch: {
      'form.name': function(){
        this.error = "";
      },
    },
    methods: {
      ...mapActions({ 
        rename: "rename"
      }),
      hideModal(){
        this.error = null;
        this.$refs['modal-rename'].hide();
      },
      async rename(){
        this.form.type = this.selection[0].type;
        this.form.uid = this.selection[0].uid;

        if(this.getElements.filter(item => item.type == this.form.type).find(item => item.name == this.form.newName)){
          this.error = "Уже есть элемент с таким именем";
        }
        else if(this.form.newName == ""){
          this.error = this.form.type == "Directory" ? "Укажите имя папки" : "Укажите имя файла";
        }else{
          await this.rename(this.form)
        }
      },
    }
}
</script>

<style>
</style>