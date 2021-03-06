import {action} from 'mobx';
import request from 'superagent';

import dataStore from '../stores/dataStore';
import fixtures from '../fixtures/apps.fixture';

class DataActions {

  constructor() {
    this.getApps()
  }

   @action getApps() {
     if(process.env.NODE_ENV === 'test') {
       dataStore.apps.clear();
       fixtures.services.map((app) => {
         dataStore.apps.push(app);
       })
       return
     }

     request
      .get(`services`)
      .end(function(err, res) {
        if (err || !res.ok) {
          return
        }

        dataStore.apps.clear();
        res.body.services.map((app) => {
          dataStore.apps.push(app);
        })
      });
   }
}

export default new DataActions;
