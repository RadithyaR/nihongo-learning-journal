"use client"

import { useEffect, useState, Suspense } from 'react';
import { useSearchParams, useRouter } from 'next/navigation';
import { api } from '@/lib/axios';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';

function VerifyEmailContent() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const token = searchParams.get('token');
  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading');
  const [message, setMessage] = useState('Verifying your email...');

  useEffect(() => {
    if (!token) {
      setStatus('error');
      setMessage('Verification token is missing.');
      return;
    }

    const verifyEmail = async () => {
      try {
        await api.post('/auth/verify-email', { token });
        setStatus('success');
        setMessage('Your email has been successfully verified. You can now log in.');
      } catch (error: any) {
        setStatus('error');
        setMessage(error.response?.data?.message || 'Failed to verify email. The link might be invalid or expired.');
      }
    };

    verifyEmail();
  }, [token]);

  return (
    <div className="flex items-center justify-center min-h-[calc(100vh-100px)] px-4">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <CardTitle className="text-2xl font-bold text-slate-800">Email Verification</CardTitle>
          <CardDescription>
            {status === 'loading' && "We're verifying your email address..."}
            {status === 'success' && "Verification successful!"}
            {status === 'error' && "Verification failed"}
          </CardDescription>
        </CardHeader>
        <CardContent className="flex flex-col items-center gap-6">
          <p className={`text-center ${status === 'error' ? 'text-red-500' : 'text-slate-600'}`}>
            {message}
          </p>
          
          {status !== 'loading' && (
            <Button 
              className="w-full" 
              onClick={() => router.push('/login')}
            >
              Go to Login
            </Button>
          )}
        </CardContent>
      </Card>
    </div>
  );
}

export default function VerifyEmailPage() {
  return (
    <Suspense fallback={<div className="flex justify-center mt-20">Loading...</div>}>
      <VerifyEmailContent />
    </Suspense>
  );
}
