<template>
    <div v-if="errorResponse" class="alert alert-danger" role="alert">
    <strong>Error {{ errorResponse.code }}: </strong>{{ errorResponse.msg }}
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue} from 'vue-facing-decorator';
import type { ErrorResponse } from '@/interfaces/error.interface';

@Component
export default class ErrorView extends Vue {
    @Prop()
    errorResponse !: ErrorResponse;


}
</script>