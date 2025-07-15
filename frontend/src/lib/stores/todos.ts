import { writable, type Writable } from 'svelte/store';
import type { Todo } from '$lib/types';

export const todos: Writable<Todo[]> = writable([]);
