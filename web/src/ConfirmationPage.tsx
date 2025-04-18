import { useNavigate, useParams } from 'react-router-dom';
import { API_URL } from './const';

export function ConfirmationPage() {
  const { token = '' } = useParams();
  const redirect = useNavigate();

  const handleConfirm = async () => {
    const resp = await fetch(`${API_URL}/users/activate/${token}`, {
      method: 'PUT',
    });

    if (resp.ok) {
      redirect('/');
    } else {
      console.error('Failed to activate user');
    }
  };

  return (
    <div>
      <h1>Confirmation</h1>
      <button onClick={handleConfirm}>Click to confirm</button>
    </div>
  );
}
