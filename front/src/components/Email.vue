<template>
    <div v-if="mail">
        <div class="col py-3 subject">
            {{ mail['Subject'] }}
        </div>
        <div class="col py-3 message">
            {{ mail['message'] }}
        </div>
        <div class="col py-2 date">
            <pre>Date: {{ formatDate(mail['Date']) }}</pre>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-facing-decorator';

@Component
export default class Email extends Vue {
    
    @Prop({
        required: true
    })
    mail: any;

    formatDate(date: string) {
        return date.split('-')[0];
    }

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