<template>
  
  <ErrorView v-if="error" :errorResponse="error"/>

  <Spinner v-if="loading" />

  <table v-if="mails" class="table">
    <thead class="table-dark">
      <tr>
        <th scope="col">Subject</th>
        <th scope="col">From</th>
        <th scope="col">To</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="mail in mails" :key="mail['_id']" v-on:click="getSelectedMail(mail['_id'])">
        <td>{{ mail['_source']['Subject'] }}</td>
        <td>{{ mail['_source']['From'] }}</td>
        <td>{{ mail['_source']['To'] }}</td>
      </tr>
    </tbody>
  </table>

  <nav aria-label="Page navigation">
    <ul class="pagination">
      <li class="page-item">
        <a class="page-link" @click="getMails()" aria-label="Previous">
          <span aria-hidden="true">&laquo;</span>
          <span class="sr-only">Previous</span>
        </a>
      </li>
      <li class="page-item">
        <a class="page-link" @click="getMails()" aria-label="Next">
          <span aria-hidden="true">&raquo;</span>
          <span class="sr-only">Next</span>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch, Emit } from 'vue-facing-decorator';
import axios from 'axios';

import Spinner from '@/components/Spinner.vue';
import ErrorView from './Error.vue';

import type { ErrorResponse } from '@/interfaces/error.interface';


@Component({
  components: {
    Spinner,
    ErrorView
  }
})
export default class Tabla extends Vue {
  
  @Prop()
  search!: string;

  error !: ErrorResponse;
  loading : boolean = false;
  

  url: string = 'http://localhost:3000/api';
  skip: number = 0;
  limit: number = 10;
  mails!: [];

  mail!: any;

  beforeMount() {
    this.getLastMails();
  }

  @Emit()
  change() {
    return this.mail;
  }

  @Watch('search')
  searchChange() {
    this.getMails();
  }

  getSelectedMail(id: string) {
    const mailData = this.mails.find(mail => mail['_id'] == id);
    this.mail = mailData!['_source'];
    this.change();
  }

  getLastMails() {
    this.loading = true;
    axios.get(`${this.url}/${this.skip}-${this.limit}`)
      .then((response: any) => {
        const { hits } = response.data;
        this.mails = hits.hits;
      })
      .catch((err: any) => {
        if (err.response) {
          this.error = {
            code: err.response.status,
            msg: err.response.data
          }
        } else if (err.request) {
          this.error = {
            msg: 'No se pudo conectar con el servicio. Verifique su conexión a internet.'
          }
        }
      })
      .then( () => {
        this.loading = false;
      })
  };

  getMails() {
    this.loading = true;
    this.mails = [];
    axios.get(`${this.url}/search/${this.search}-${this.skip}-${this.limit}`)
      .then((response: any) => {
        this.skip += this.limit;
        const { hits } = response.data;
        this.mails = hits.hits;
      })
      .catch((err: any) => {
        if (err.response) {
          this.error = {
            code: err.response.status,
            msg: err.response.data
          }
        } else if (err.request) {
          this.error = {
            msg: 'No se pudo conectar con el servicio. Verifique su conexión a internet.'
          }
        }
      })
      .then( () => {
        this.loading = false;
      });
  }

}

</script>

