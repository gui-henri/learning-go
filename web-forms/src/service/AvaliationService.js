export const AvaliationService = {
    async getFormData(avaliationFormStore) {
        try {
            const currentVersion = avaliationFormStore.avaliationForm?.Version || 0;

            const response = await fetch("http://localhost:8090/Avaliation/form", {
                method: 'GET',
                headers: {
                    'If-None-Match': currentVersion.toString()
                }
            })

            if (response.status === 304) {
                console.log("Formulário não modificado (304). Usando cache local.");
                return avaliationFormStore.avaliationForm;
            }

            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }

            const result = await response.json();
            avaliationFormStore.avaliationForm = result;

            return result;

        } catch (error) {
            console.error('Failed to get form:', error);
            alert('Erro ao baixar o formulário.');
        }
    },
    async submitFormData(payload) {
        try {
            console.log(JSON.stringify(payload))
            const response = await fetch('http://localhost:8090/Avaliation', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }

            const result = await response.json();
            console.log('Success:', result);
            alert('Formulário salvo com sucesso!');

        } catch (error) {
            console.error('Failed to save form:', error);
            alert('Erro ao salvar o formulário.');
        }
    }
}