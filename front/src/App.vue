<template>

  <NavBar />
  <div class="container pt-5">
    <div class="col">
      <InputForm @change="onChangeSearch" />
    </div>
    <div class="d-flex justify-content-between">
      <div class="col">
        <Tabla @change="mailSelected" :search="search" />
      </div>
      <div class="col m-5">
        <Email :mail="mail"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-facing-decorator';

import NavBar from '@/components/NavBar.vue';
import InputForm from '@/components/InputForm.vue';
import Tabla from '@/components/Tabla.vue';
import Email from '@/components/Email.vue';

@Component({
  components: {
    NavBar,
    InputForm,
    Tabla,
    Email
  }
})
export default class AppVue extends Vue {
  search!: string;
  mail !: any;

  public onChangeSearch(value: string): void {
    this.search = value;
  }

  public mailSelected(mail: any): void {
    this.mail = mail;
  }
}
</script>