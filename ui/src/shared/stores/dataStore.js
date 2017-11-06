import { observable } from "mobx";

class DataStore {
  @observable apps = [];
  @observable filter = '';
}

let store = new DataStore;

export default store;
