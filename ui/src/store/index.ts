import Vue from 'vue';
import Vuex from 'vuex';
import { getStoreBuilder } from 'vuex-typex';

Vue.use(Vuex);

export interface RootState {}

const store = getStoreBuilder<RootState>().vuexStore({});

export default store;
