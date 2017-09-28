import React from 'react';
import {Route, IndexRoute} from 'react-router';

import App from './components/App'
import UserService from './containers/UserService'
import About from './components/About'
// <IndexRoute  component={SSMControl}/>

export default (
  <Route path={'/'} component={App}>
    <Route path={'about'} component={About}/>
    <IndexRoute  component={UserService}/>
  </Route>
)
