<template>

  <nav class="navbar navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">
        <img src="./assets/logo.svg" alt="Logo" width="30" height="24" class="d-inline-block align-text-top">
        Test
      </a>
    </div>
  </nav>

  <div class="container pt-5">
    <div class="col">
      <div class="input-group mb-3">
        <input v-model="search" class="form-control" placeholder="Search" aria-label="Search" aria-describedby="search">

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
            <tr v-for="mail in mails" :key="mail['_id']"
                v-on:click="ping(mail['_id'])">
              <td>{{ mail['_source']['Subject'] }}</td>
              <td>{{ mail['_source']['From'] }}</td>
              <td>{{ mail['_source']['To'] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="col m-5">
        <div v-if="mail">
          <div class="col"><p>{{mail['Subject']}}</p></div>
          <div class="col">{{mail['message']}}</div>
          <div class="col">{{mail['Date']}}</div>
        
        </div>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { Observable, catchError, pipe, debounce, Subject, debounceTime } from 'rxjs';
import axios from 'axios';


export default {
  name: 'MailsList',
  data() {
    return {
      mails: [],
      mail: undefined,
      total: 0,
      errorMsg: '',
      search: '',
      URL: 'http://localhost:3000/api',
      debouncer: new Subject<String>()
    }
  },
  methods: {
    getMails() {
      axios.get(`${this.URL!}/search/${this.search}-0-10`)
        .then((response: any) => {
          const { hits } = response.data;
          this.mails = hits.hits;
          this.total = hits.total;
        })
        .catch((err: any) => {
          this.errorMsg = 'Error: ' + err;
        })
    },
    ping(id: string) {
      const mailData = this.mails.find( mail => mail['_id'] === id);
      this.mail = mailData!['_source'];
      console.log(this.mail)
    },
    debounceSearch() {
      this.debouncer
        .pipe(debounceTime(300))
        .subscribe(valor => {
          console.log(valor)
        })
    }
  }
}
</script>


<style>

</style>
