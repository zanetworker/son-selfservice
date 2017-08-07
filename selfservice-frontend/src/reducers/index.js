import {combineReducers} from 'redux';
import fsmsReducer from './fsms'

const rootReducer = combineReducers({
  fsms: fsmsReducer,
});

export default rootReducer;
