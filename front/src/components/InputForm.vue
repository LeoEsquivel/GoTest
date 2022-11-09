<template>
    <div class="input-group mb-3">
        <input v-on:keydown.enter="change()" 
               v-model="search" type="text" 
               class="form-control" 
               placeholder="Search">

        <span class="input-group-text" id="search">
            <a @click="change()">
                <span class="material-symbols-outlined">
                    search
                </span>
            </a>
        </span>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Emit } from 'vue-facing-decorator';

@Component
export default class InputForm extends Vue {
    search!: string;

    @Emit()
    public change() {
        return this.search;
    }

}

</script>