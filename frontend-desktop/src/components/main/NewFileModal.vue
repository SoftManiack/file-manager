<template>
    <b-modal 
        content-class="l-main_modal modal"
        header-class="modal_header"
        id="modal-newfile"
        ref="modal-newfile"
        hide-footer 
        centered>
        <div class="modal_body">
          <h5 class="mb-4">Новый тектовый файл</h5>
          <b-row class="px-3">
            <b-form-input 
              class="modal_input"  
              :class="{'modal_input--error': Error}"
              placeholder=".txt"
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
            <button class="modal_create ml-2" size="sm" @click="createTextFile">Создать</button>
          </b-row>
        </div>
    </b-modal>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
export default {
    name: "NewFileModal",
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
        this.error = "";
      },
    },
    methods: {
      ...mapActions({ 
        createFile: "createFile"
      }),
      hideModal(){
        this.error = null;
        this.$refs['modal-newfile'].hide();
      },
      async createTextFile(){
        if(this.getElements.filter(item => item.type == "File").find(item => item.name == this.form.name)){
          this.error = "Уже есть элемент с таким именем";
        }
        else if(this.form.name == ""){
          this.error = "Укажите имя папки";
        }else{
          this.form.rootUid = this.rootUid;
          this.error = null;
          await this.createFile(this.form);
          this.hideModal()
        }
      },
    }
}
</script>

<style>
</style>