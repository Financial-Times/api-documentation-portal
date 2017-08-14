import React from "react";
import ReactDOM from "react-dom";
import { Provider } from 'mobx-react';
import { HashRouter, Route } from 'react-router-dom'

import store from '../shared/stores/dataStore';
import '../shared/actions/dataActions';

import Main from '../shared/main';
import Documentation from '../shared/documentation';

// import styles
import "./styles/main.scss";

ReactDOM.render(
   <Provider store={store}>
      <HashRouter>
         <div>
            <Route exact path="/"
               component={Main}
            />
            <Route path="/docs/:name"
               component={Documentation}
            />
         </div>
      </HashRouter>
   </Provider>,
  document.getElementById('app')
);
