import {combineReducers} from 'redux';
import fsmsReducer from './fsms'
import modalReducer from './modal'

const rootReducer = combineReducers({
  fsms: fsmsReducer,
  modal: modalReducer,
});

export default rootReducer;
