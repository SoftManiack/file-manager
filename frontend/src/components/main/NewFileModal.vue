<template>
    <b-modal 
        content-class="l-main_new-modal-fd new-modal-fd"
        header-class="new-modal-fd_header"
        id="modal-2"
        ref="modal-2"
        hide-footer 
        centered>
        <div class="new-modal-fd_body">
          <h5 class="mb-4">Новый тектовый файл</h5>
          <b-row class="px-3">
            <b-form-input 
              class="new-modal-fd_input"  
              :class="{'new-modal-fd_input--error': Error}"
              placeholder=".txt"
              v-model="form.name"
            >
            </b-form-input>
            <transition name="fade">
              <div v-if="Error" class="new-modal-fd_error mt-2">
                {{ Error }}
              </div>
            </transition>
          </b-row>
          <b-row class="mt-3 px-3" align-h="end">
            <button class="new-modal-fd_exit" size="sm" @click="hideModal" >Отмена</button>
            <button class="new-modal-fd_create ml-2" size="sm" @click="createTextFile">Создать</button>
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
        this.$refs['modal-2'].hide();
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