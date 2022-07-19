<template>
    <b-modal 
      content-class="l-main_modal modal"
      header-class="modal_header"
      id="modal-newfolder"
      ref="modal-newfolder"
      hide-footer 
      centered>
      <div class="modal_body">
        <h5 class="mb-4">Новая папка</h5>
        <b-row class="px-3">
          <b-form-input 
            class="modal_input"  
            :class="{'modal_input--error': Error}"
            placeholder="Имя папки"
            v-model="form.name"
          >
          </b-form-input>
          <transition name="fade">
            <div v-if="Error" class="modal_error mt-2">
              {{ Error }}
            </div>
          </transition>
        </b-row>
        <b-row class="mt-3 px-3" align-h="end">
          <button class="modal_exit" size="sm" @click="hideModal" >Отмена</button>
          <button class="modal_create ml-2" size="sm" @click="createDirectory">Создать</button>
        </b-row>
      </div>
    </b-modal>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "NewFolderModal",
    data(){
      return {
        form: { 
          name: "",
          rootUid: "",
        },
        error: "",
      }
    },
    emits: ['closeMenu'],
    computed: {
      ...mapGetters({
        getElements: "elements",
        rootUid: "rootUid",
      }),
      Error(){
        return this.error
      }
    },
    watch: {
      'form.name': function(){
        this.error = ""
      },
    },
    methods: {
      ...mapActions({ 
        createDir: "createDirectory"
      }),
      hideModal(){
        this.error = null;
        this.$refs['modal-newfolder'].hide();
      },
      async createDirectory(){
        if(this.getElements != null && this.getElements.find(item => item.name == this.form.name)){
          this.error = "Уже есть элемент с таким именем";
        }
        else if(this.form.name == ""){
          this.error = "Укажите имя папки";
        }else{
          alert(this.rootUid)
          this.form.rootUid = this.rootUid;
          this.error = null;
          await this.createDir(this.form);
          this.hideModal();
        }
      },
    }
}
</script>

<style>
</style>