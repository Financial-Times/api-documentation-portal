import React from "react";
import ReactDOM from "react-dom";
import { Provider } from 'mobx-react';
import { HashRouter, Route } from 'react-router-dom'

import store from '../shared/stores/dataStore';
import '../shared/actions/dataActions';

import Documentation from '../shared/documentation';

// import styles
import "./styles/vendor/_reset.scss";

ReactDOM.render(
   <Provider store={store}>
      <HashRouter>
         <div>
            <Route path="/docs/:name"
               component={Documentation}
            />
         </div>
      </HashRouter>
   </Provider>,
  document.getElementById('app')
);
