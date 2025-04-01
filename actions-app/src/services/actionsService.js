import api from './api';

export const getActions = () => api.get('/actions');
export const getBestActions = () => api.get('/actions/best-actions');