// src/composables/useStepAccordion.js
import { ref, watch, nextTick } from 'vue';

export function useStepAccordion(props, emit) {
    const internalIndex = ref(null);

    watch(() => props.isActive, (isOpen) => {
        if (isOpen) {
            internalIndex.value = 0; // Open the tab
        } else {
            internalIndex.value = null; // Close the tab
        }
    }, { immediate: true });

    // Helper to emit the event to the parent
    const nextStep = () => {
        emit('next-step');
    };

    return {
        internalIndex,
        nextStep
    };
}