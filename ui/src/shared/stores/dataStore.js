import { observable } from "mobx";

class DataStore {
  @observable apps = [];
}

let store = new DataStore;

export default store;
