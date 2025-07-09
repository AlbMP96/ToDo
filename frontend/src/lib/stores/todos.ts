import { writable } from 'svelte/store';
import type { Todo } from '$lib/types';

export const todos = writable<Todo[]>([]);
