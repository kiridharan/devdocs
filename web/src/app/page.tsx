import Navbar from '../components/Navbar';
import Hero from '../components/Hero';
import ProblemSolution from '../components/ProblemSolution';
import ValueProp from '../components/ValueProp';
import HowItWorks from '../components/HowItWorks';
import FeatureBenefit from '../components/FeatureBenefit';
import SocialProof from '../components/SocialProof';
import FAQ from '../components/FAQ';
import Footer from '../components/Footer';

export default function Home() {
  return (
    <main className="min-h-screen flex flex-col">
      <Navbar />
      <Hero />
      <ProblemSolution />
      <ValueProp />
      <HowItWorks />
      <FeatureBenefit />
      <SocialProof />
      <FAQ />
      <Footer />
    </main>
  )
}
