<template>
  <b-container class="l-main" fluid>
    <b-row class="l-main_header c-header pt-2">
        <b-col class="d-flex" md="7">
          <b-breadcrumb :items="items"></b-breadcrumb>
        </b-col>
        <b-col class="c-header_buttons d-flex mt-2" md="4" align-h="end" align-self="end">
          <b-row>
            <button id="popover-2">
              <svg width="23" height="22"><use xlink:href="@/assets/icons/sprite.svg#add"></use></svg>
              Новая папка
            </button> 
            <label class="c-header_upload" for="input_upload">    
              <svg for="input_upload" width="23" height="22"><use xlink:href="@/assets/icons/sprite.svg#download-arrow"></use></svg>
              Загрузить...       
              <input class="c-header_upload" name="input_upload" type="file" ref="file" v-on:change="fileUpload()"/>
            </label>
          </b-row>
        </b-col>
        <b-col class="d-flex flex-row-reverse">
          <b-row class="mt-2 mr-1 c-header_view d-flex justify-content-center" align-h="end">
            <div class="c-header_wrapper d-flex justify-content-center mr-2" @click="switchList" :class="{'c-header_wrapper--active': viewList}">
              <svg width="16" height="14" :fill ="activeList"><use xlink:href="@/assets/icons/sprite.svg#list"></use></svg>
            </div>
            <div class="c-header_wrapper d-flex justify-content-center" @click="switchBlock" :class="{'c-header_wrapper--active': viewBlock}">
              <svg width="16" height="14" :fill ="activeGrid"><use xlink:href="@/assets/icons/sprite.svg#grid"></use></svg>
            </div>
          </b-row>
        </b-col>
      </b-row>

      <b-container 
        fluid class="l-main_data pl-3 pr-2 pb-4" 
        @click.capture.self="clearSelection"
        @click="closeMenu"
        @contextmenu.self="switchMenu"
        :class="{'l-main_data--bar':selection.length}">
        <!-- View Block -->
        <div v-if="getElements[0].uid != '' "> 
          <div @click.capture.self="clearSelection" @contextmenu.self="switchMenu" v-if="viewBlock" class="d-flex flex-row flex-wrap mt-3">
            <div @click.capture.self="clearSelection" @contextmenu.self="switchMenu" v-for="(item, key) in getElements" :key="key">
              <Block
                :uid="item.uid"
                :name="item.name"
                :size="item.size"
                :type="item.type"
                :isFavorites="item.isFavorites"
                :date="item.date"
                :extension="item.extension"
                :selection="selection"
                @selectBlock="select(item)"
                @selectManyBlock="selectMany(item)"
                @switchContextmenu="switchContextmenu($event, item.uid)"
                @openDirectory="openDirectory(item)"
              />
            </div>
          </div>

        <!-- View Table -->
        <div v-if="viewList">
          <b-row>
            <b-table class="l-main_table main-table" small :items="getElements" :fields="fields" >
              <template #cell(isFavorites)="data">
                <div class="main-table_favorite" v-if="data.item.isFavorites">
                    <svg  width="6" height="6"><use xlink:href="@/assets/icons/sprite.svg#dot"></use></svg>
                </div>
                <div class="main-table_heart main-table_favorite" v-if="!data.item.isFavorites">
                    <svg  width="13" height="13"><use xlink:href="@/assets/icons/sprite.svg#heart"></use></svg>
                </div>
              </template>
              <template #cell(name)="data">
                <div class="d-flex">
                  <div v-if="data.item.type == 'File'" class="main-table_iconfile">
                    <svg  width="17" height="17"><use xlink:href="@/assets/icons/sprite.svg#empry-file"></use></svg>
                  </div>
                  <div v-if="data.item.type == 'Directory'" class="main-table_iconfolder" >
                    <svg  width="19" height="19"><use xlink:href="@/assets/icons/sprite.svg#folder"></use></svg>
                  </div>
                  {{data.item.name}}
                </div>
              </template>
                <b-button></b-button>
            </b-table>
          </b-row>
        </div>
        </div>
        <div  
          @click="closeMenu"
          @contextmenu.self="switchMenu"
          class="l-main_empty-dir d-flex justify-content-center align-items-center" v-else 
        >
          <div>
            <svg  width="161" height="161"><use xlink:href="@/assets/icons/sprite.svg#folder-empty"></use></svg>
            <h3 class="justify-content-center"> Папка пуста  </h3>
          </div>
        </div>
      </b-container>
      
      <div v-if="selection.length">
        <SelectionBar
          :selection="selection"
        />
      </div>

      <!-- Popovers -->
      <b-popover
        target="popover-2"
        placement="bottom"
        triggers="focus"
      >
        <div class="main-popover">
          <b-row>
            <b-form-input  placeholder="Имя папки"></b-form-input>
          </b-row>
          <b-row class="mt-3 px-3" align-h="end">
            <b-button class="main-popover_btnCancel" size="sm">Отмена</b-button>
            <b-button class="main-popover_btnCreate ml-2" size="sm">Создать</b-button>
          </b-row>
        </div>
      </b-popover>

      <!--Modals-->
      <NewFileModal/> 
      <NewFolderModal/>
      <InfoFileModal :selection="selection"/>
      <InfoFolderModal :selection="selection"/>
      <RenameModal :selection="selection"/>
      
      <!--Contextmenu-->
      <MainContextmenu @closeMenu="closeMenu"/>
      <FolderContextmenu @closeMenu="closeMenu"/>
      <FileContextmenu @closeMenu="closeMenu"/>

      <!-- BlockMenu https://icons.getbootstrap.com/  -->

      
      <!-- BlockMenu -->

  </b-container>
</template>

<script>
import { contextMenu } from '@/utils/contextMenu'
import { mapGetters, mapActions } from "vuex";

export default {
    name: "Main",
    components: {
      Block: () => import("@/components/main/Block.vue"),
      SelectionBar: () => import("@/components/main/SelectionBar.vue"),
      NewFolderModal: () => import("@/components/main/NewFolderModal.vue"),
      NewFileModal: () => import("@/components/main/NewFileModal.vue"),
      InfoFileModal: () => import("@/components/main/InfoFileModal.vue"),
      InfoFolderModal: () => import("@/components/main/InfoFolderModal.vue"),
      RenameModal: () => import("@/components/main/RenameModal.vue"),
      MainContextmenu: () => import("@/components/main/MainContextmenu.vue"),
      FolderContextmenu: () => import("@/components/main/FolderContextmenu.vue"),
      FileContextmenu: () => import("@/components/main/FileContextmenu.vue"),
    },
    data() {
      return {
        viewList: false,
        viewBlock: true,
        file: '',
        items: [
          {
            text: 'Admin',
            href: '#'
          },
          {
            text: 'Manage',
            href: '#'
          },
          {
            text: 'Library',
            active: true
          },
        ],
        selection: [],
        fields : [
          {
            key: 'isFavorites',
            sortable: true,
            label: "Избранное"
          },
          {
            key: 'name',
            sortable: true,
            label: 'Имя'
          },
          {
            key: 'size',
            sortable: true,
            label: 'Размер'
          },
          {
            key: 'type',
            sortable: true,
            label: 'Тип'
          },
          {
            key: 'date',
            sortable: true,
            label: 'Дата'
          }
        ],
      }
    },
    computed: {
      ...mapGetters({
        getElements: "elements",
        getRootUid: "rootUid",
        getLoadingStatus: "getLoadingStatus"
      }),
      activeGrid(){
        return this.viewBlock ? '#000000': '#CCCCCC';
      },
      activeList(){
        return this.viewList ? '#000000': '#CCCCCC';
      }
    },
    methods: {
      ...mapActions({ loadElements: "getElements"}),
      switchBlock(){  
        this.viewBlock = true;
        this.viewList = false;
      },
      switchList(){
        this.viewList = true;
        this.viewBlock = false;
      },
      select(item){
          this.selection.length = 0;
          this.selection.push(item);
      },
      selectMany(item){
        this.selection.push(item);
      },
      clearSelection(){
        this.selection = [];
      },
      fileUpload(){
        this.file = this.$refs.file.files[0]
        let formData = new FormData();
        formData.append('file', this.file);
      },
      async openDirectory(item){
        if(item.type == 'Directory'){
          await this.loadElements(item.uid)
          this.$router.push(`/v1/fm/${this.getRootUid}`);
        }
      },
      switchMenu(event){
        this.closeMenu();
        event.preventDefault();
        contextMenu("contextmenu-1", event);
      },
      switchContextmenu(event, uid){
        this.closeMenu();
        event.preventDefault();
        
        if(this.getElements.find(item => item.uid == uid).type == 'Directory'){
          contextMenu("contextmenu-2", event);
        }
        else{
          contextMenu("contextmenu-3", event)
        }
      },
      closeMenu(){
        var contextElement = document.getElementById("contextmenu-1");
        contextElement.style.display = 'none';
        contextElement = document.getElementById("contextmenu-2");
        contextElement.style.display = 'none';
        contextElement = document.getElementById("contextmenu-3");
        contextElement.style.display = 'none';
      }
    },
}
</script>

<style>
</style>