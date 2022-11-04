<template>

  <nav class="navbar navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">
        <img src="./assets/logo.svg" alt="Logo" width="30" height="24" class="d-inline-block align-text-top">
        Test
      </a>
    </div>
  </nav>

  <div>
    <TablaVue />
  </div>

  <div class="container pt-5">
    <div class="col">
      <div class="input-group mb-3">
        <input v-on:keydown.enter="getMails()"
               v-model="search" 
               class="form-control" 
               placeholder="Search" 
               aria-label="Search" 
               aria-describedby="search">

        <span class="input-group-text" id="search">
          <a @click="getMails()">
            <span class="material-symbols-outlined">
              search
            </span>
          </a>
        </span>
      </div>
    </div>
    <h3 v-if="errorMsg">{{ errorMsg }}</h3>
    <div class="d-flex justify-content-between">
      <div class="col">
        <table class="table">
          <thead class="table-dark">
            <tr>
              <th scope="col">Subject</th>
              <th scope="col">From</th>
              <th scope="col">To</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="mail in mails" :key="mail['_id']" v-on:click="showMail(mail['_id'])">
              <td>{{ mail['_source']['Subject'] }}</td>
              <td>{{ mail['_source']['From'] }}</td>
              <td>{{ mail['_source']['To'] }}</td>
            </tr>
          </tbody>
        </table>

        <nav  v-if="search"
              aria-label="Page navigation">
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
      </div>
      <div class="col m-5">
        <div v-if="mail">
          <div class="col py-3 subject">
            {{mail['Subject']}}
          </div>
          <div class="col py-3 message">
            {{ mail['message'] }}
          </div>
          <div class="col py-2 date">
            <pre>Date:{{ formatDate(mail['Date']) }}</pre>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { Observable, catchError, pipe, debounce, Subject, debounceTime } from 'rxjs';
import axios from 'axios';

// import TablaVue from './views/Tabla.vue';

export default {
  name: 'MailsList',
  data() {
    return {
      mails: [],
      mail: undefined,
      total: 0,
      errorMsg: '',
      search: '',
      skip: 0,
      limit: 10,
      URL: 'http://localhost:3000/api',
      debouncer: new Subject<String>()
    }
  },
  methods: {
    getLastMails() {
      axios.get(`${this.URL!}/${this.search}-${this.skip}-${this.limit}`)
        .then((response: any) => {
          this.skip += this.limit
          const { hits } = response.data;
          this.mails = hits.hits;
          this.total = hits.total;
        })
        .catch((err: any) => {
          this.errorMsg = 'Error: ' + err;
        })
    },
    getMails() {
      axios.get(`${this.URL!}/search/${this.search}-${this.skip}-${this.limit}`)
        .then((response: any) => {
          this.skip += this.limit
          const { hits } = response.data;
          this.mails = hits.hits;
          this.total = hits.total;
        })
        .catch((err: any) => {
          this.errorMsg = 'Error: ' + err;
        })
    },
    showMail(id: string) {
      const mailData = this.mails.find(mail => mail['_id'] === id);
      this.mail = mailData!['_source'];
    },
    formatDate(dateString: string) {
      const date  = dateString.split('-');
      return date[0]
    },
    debounceSearch() {
      this.debouncer
        .pipe(debounceTime(300))
        .subscribe(valor => {
          console.log(valor)
        })
    }
  },
  beforeMount() {
      this.getLastMails();
  },
}
</script>


<style>

  .subject {
    color: black;
    font-weight: bold;
    text-transform: capitalize;
    font-size: 1.3rem;
  }

  .message {
    color: black;
    text-align: justify;
  }

</style>
