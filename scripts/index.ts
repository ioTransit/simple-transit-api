import Alpine from 'alpinejs'
import { createIcons } from './lucide.js';
import * as htmx from 'htmx.org';

const hello = () => {
  console.log('hello');
  createIcons()
  // @ts-expect-error window update
  window.Alpine = Alpine
  // @ts-expect-error window update
  window.htmx = htmx
};

hello();
