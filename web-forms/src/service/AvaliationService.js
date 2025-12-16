import { useAvaliationForm } from "@/store/evaluation/form";

const avaliationForm = useAvaliationForm()

export const AvaliationService = {
    async getFormData() {
        try {
            const currentVersion = avaliationForm.avaliationForm?.Version || 0;

            const response = await fetch("http://localhost:8090/Avaliation/form", {
                method: 'GET',
                headers: {
                    'If-None-Match': currentVersion.toString()
                }
            })

            if (response.status === 304) {
                console.log("Formulário não modificado (304). Usando cache local.");
                return avaliationForm.avaliationForm;
            }

            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }

            const result = await response.json();
            avaliationForm.avaliationForm = result;

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
            avaliationForm.avaliationId = result.id;
            console.log('Success:', result);

            alert('Formulário salvo com sucesso!');

        } catch (error) {
            throw error;
        }
    },
    async appendToAvaliation(id, payload) {
        try {
            console.log(JSON.stringify(payload))

            const response = await fetch(`http://localhost:8090/Avaliation/${id}`, {
                method: 'PUT',
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
        } catch (error) {
            console.error('Failed to save form:', error);
            alert('Erro ao salvar o formulário.');
            throw error;
        }
    }
}