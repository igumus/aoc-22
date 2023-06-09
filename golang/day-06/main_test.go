package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay06(t *testing.T) {
	testcases := []struct {
		input    string
		f        func(string, int) int
		capacity int
		result   int
	}{
		{
			f:        findDistinct,
			capacity: 4,
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			result:   7,
		},
		{
			f:        findDistinct,
			capacity: 4,
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			result:   5,
		},
		{
			f:        findDistinct,
			capacity: 4,
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			result:   6,
		},
		{
			f:        findDistinct,
			capacity: 4,
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			result:   10,
		},
		{
			f:        findDistinct,
			capacity: 4,
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			result:   11,
		},
		{
			f:        findDistinct,
			capacity: 4,
			input:    "stftmtvvtvqqczqqnjnwwlqqdzdnnsvnsswbbwsstvvssfjsjbjfjmjpjzpplpppjzjqqdzzhqqqqtcccbzzzwzrrrdqdldpdsppmqmmnwwjddnqqscclncllvhllqpllchhbccfcbcgbcgcfcncsnstsddldzldlmljjfbjbzbccmrmrppqmqsswbwqwdwwcnwwhrhppfsfvsvrrfllhglhlggjpggzjgzggnvvqfvvhffpwpmwpmmwvmvrmrbmbzmzbbvgbbcfbcfbfppnzpzrrszzqgzgjgddmdwmwrmwmzznqzqhqhvvsslppsrrljjfpfcpfpbbrjjwjmjpmpfmfzfvzfftptzzbmzmddpvdddqmmzjzbbhmmwqmmmbgmmttrhrqrvqvzvdvzdvvmsvmmqlmmtddvlvttrtvtvcttvssnwwbccqmmgbbqrqlqjllmslsmslltrtffzfpfzpffvwvffsllgvgtgwtwnttfzznzqzztfzfvzznnwzzcvcqvvdwwsnsvnnthhnphpssmjmfmhfmhmgmllmsmrsrrmmhsswjwqqdbbghhpsptsswvsvfvcffqcchlhfhvhjhdjhddvjjpmmsrsqrrngrrmvmsvsllrmlrlprpggqzqmmvlvvwrwnrwwzztrzrbrdbrdbbhlhchjchhtthpthphplhhlphpjjsddtppvbpbbmnmgnmmqbbrhhfrfpfjpjgjljcctwtmwtwvvvmsvmmnjmnjmnmqmvvggtzzctzttszsvvjwjqjmjbjvbbshbhghlghhpvpnvvqmmgjmmggqvqtvqtvqtvvmlljbbhdhshnnwqwbbrnnwswmwfmfttjztjzjsjzjdjbdjjzfzdzgzhghgpgqqdzdhhnwnjnnhrnrqqjsqsllbzzzcnzzmjmvmrvrgrfgrrqmrmpmpzmzqqfwqfwfjjhphgpptzzwmmfjjvbjvbbqsbsbcbppjzpjjzpjpvjjzdjjgcgppvrvjrjpjspjppntpnnjlnlrnrdndjnnsbnnppvspprbrddjrdrmrfmmtmtbtntftjjsscvcbvvcnccnbnrbbmlmddhpdptddmrmmjddnjjtztctqqhhttqctqtbbnfnwnmndnzzljlgltgltljtjllwjjbrjbrbpbgppmzpmpggdnnmznnhhmfhhssgtgvvlsljltlppdqqbtbvbsvsswbsbhsbsmmwswbbbbhvbvcbvvsvcsvvmfmlmgmjgmmbqbrbsrrgrsggldlmlttpbpmmtptctqcqvccnhnmnssmvsmmddlclpccfwfdfqfvqfvfqqwmmwrmmbzznpnllfttgcttnbtnnsswttfppnddpdpsphhpttnrtrlrsrmrnmmzhmhnmnwwddmhdhqqhbqqdcclnlfnnzfnzzqfqddczcqcrrjsrrswslltfltffjhhqqfcchqcqpplbltbblbcbncnpccdffcwwqzwqqlwwtjjpbbjvvljvjgvvrsszfsfcssqsvvwsvvdfvdfdpffztzrtzthzhhmwmzwzssjlsldsldsldlvvjdvvnbvnvmvmccvfcvcfcbchhzqqhgqqcbbhdbdpddwcdwwzszsfszzgbgdgzgvgsgnssfqsqjssrppsgppmmpzmppglgqlqqpzpqqfzzzrnzrnrqqzzfwzzffjfzzmnznpznpprhhvcvvvfddcrdrsscnnsznzszbznntwntnrnbnzzrnnlwlggnzzwppmbmrrncnlclzltlblhhlvhvphhsdhsdsqddgzgvvshvvzzbvbrrlplqpptplljcljlvjlvlqvvndnsdddmvddzhzjjgpjpddppzllwtwfwgfgpffmhhzllsttmqqhjhhpvpgpfjsgscnwjmwmtmptwlpfjljwgpgntrlpjfgbjqmcpzgfhrwmznqnsbpptbdrzmdtvvtdqjgrjzlphndhmlchvddglqnqsjqrfqslprsvlqjwqnsmsznptsstpvdntpttslpmqqbsdlqwpjqnzmpblgqmjrvqwsncnzdszgfsghddlnwhwzpgtddgstttvrjfjwwfrgsdjjngljqlqcrzlgsmwngbzvmjwtnqdqcgwmfhsztgrtvvfzbtstmdbqpntdpsszjthqvpbdwswfzvmrcpbgbgdmldfhvdpfsmdzfhwsrpcglsztdwqgbqszcqtqjhgntzvttldqsffftzmllptzhmhpmfsgcchfrnrchnsgcfjbgrqmvrmmhnmlnwtgwhznqfgwnlrqlpjrvfrgzcjwncvlwhpclfzngbgvmrmlzngmqlvvwhbpjzlclgrcnnvnlppqhvlrnpzvmtsbdpfbwgffgzfwvltcfvfdcnfhwvcvclwwbmshhmpgrzgltwjmqczpqzdwfjpqhmwqhvvgnpgtwrjrgwvhthtdrdpnwpbmwstgblwmbfvlwflqmfbcsgwstwvncwfcsmrpcfrrvmlbqhdtdswswfnzhgzlngwsrtlzfcgdppmjnghfrgbdqhqmslhcqddjvsslsjwqqznttqjzdlghnsvqqtwrpfbzjgwnrhhvlnbqmnvcpblzgbzltnrhzpdwvbqbtmctbzgsdjfzswrbqbzgvwjlwtmgcllnmnwcljbhbplpvtgpgjftfrbgpgmhghnjcgjfqmsbqhbgtzbfzgwmfdsgfgmgzsbgdrszfhjttbvcqjzjgbqfgswlmrrhnmnfrptvjtlnvplgznsljzfzmhghlsccnqzflfnmbhshfprhclmtfptcmtnhrjgnngnqnczvcgzzlntftjsbgpgwzbnrhzzfqmznqnzfrvjzsmtpjbswzjlbgpfftzrzbfgdblpwscbqjfrfmfnhhlhjprtlzzvwnwzsnqhnmgwsdprnrblgbclzhthftqzdljspwdzwmwhfmdzmlvqsngppdfsjdprbrhffcvcvzztjjqcffwbrvpzvzfzhjvsfvsnrmjvqmjtrjbmbqsdtjgvtbbzzfmnmrrflgcdtljpmpvqvdbzgbmhjgccgdtplllctzqpfqnsztbwdbmqgfzrcddtmwrgmwsghcfpgqssdjrtqhtjfbpjvjdnzgvpzrhbrhrhcmpbglbbrvdltdpsrwjbjzftccwgnqmnlqlpjwrfdvmlgvgqznlvsmpzsmgjstvqbqpprzlsdndpfbmqcrfgvcfvlfhmpfnnqlcqlnbbgcrrrhbtzwnwmfrnrzvgmqlmqnmnzbwflwzcmncphjqztlrzvpztqhptmfsrppvvlzcfnlrwptgccsjjjscjcwnzssmbcvtzhnscgsbrchbqbrtdzllfvmqfwznfzpzmbfwcdsfhdlddnfbdgqbqjqzdtppshwcvvcjqstdgtgbhmlqlrfrhbvfsszsmbldmwfnfgnjptdslnzwcjgmvbnqcfzjmrslrlllplbhpjcnvzmvrfzwqhmbnrvpqnvcfncgfwqlvcwpwwljssfmswctcmhgtphvjdpfnnzznfbdjcsndrczlgdjhrnrltsgbtqmqcbwlthwcgsbvqbnntcznczpmlmblwdrlmzqdztwsgjthjwtfcpgwbczmdmhttzcwvdzhdfldmwnbnfdcjgvjhrfltjnjhqhzmzrlbncwdnlgshmqhpgsdwbvmvjsfgvgqzqjqdzqzmfmrncfdgrqfrnstvpqwtltnhgrmhgmmnwvlnsfmmbjrdmrnlgfgpqncdpqgvjltpghbgffdppdcfhdqhtvrdnfvcttlrqppfnqtmzpfgsnmjqfrgtbbdzzrccsrzgfjndlrnzqmjjtgldglcpzcrhwfplvdndjtzbpfbbbpfljqnlvrdzbrvczzwdrvzlmslbjsqgqdrltmhwcdpldqldlpctcbjmsvdwlfspzlpgrhdrvjtprcwrmszvzwmmjgvsbfcztmrdgrgshgtggpzszgqhwmhdzjmzsgqsfmqvcgqqwgprgvvqlbfhpwsfbjdqtcfnfhsgzthzhbpwggbnscqbnvcdprsbgllsrdcclqggfgfdrpqlqljfwpmzdhthpczcsqrrm",
			result:   1896,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			result:   26,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			result:   29,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			result:   23,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			result:   23,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			result:   19,
		},
		{
			f:        findDistinct,
			capacity: 14,
			input:    "stftmtvvtvqqczqqnjnwwlqqdzdnnsvnsswbbwsstvvssfjsjbjfjmjpjzpplpppjzjqqdzzhqqqqtcccbzzzwzrrrdqdldpdsppmqmmnwwjddnqqscclncllvhllqpllchhbccfcbcgbcgcfcncsnstsddldzldlmljjfbjbzbccmrmrppqmqsswbwqwdwwcnwwhrhppfsfvsvrrfllhglhlggjpggzjgzggnvvqfvvhffpwpmwpmmwvmvrmrbmbzmzbbvgbbcfbcfbfppnzpzrrszzqgzgjgddmdwmwrmwmzznqzqhqhvvsslppsrrljjfpfcpfpbbrjjwjmjpmpfmfzfvzfftptzzbmzmddpvdddqmmzjzbbhmmwqmmmbgmmttrhrqrvqvzvdvzdvvmsvmmqlmmtddvlvttrtvtvcttvssnwwbccqmmgbbqrqlqjllmslsmslltrtffzfpfzpffvwvffsllgvgtgwtwnttfzznzqzztfzfvzznnwzzcvcqvvdwwsnsvnnthhnphpssmjmfmhfmhmgmllmsmrsrrmmhsswjwqqdbbghhpsptsswvsvfvcffqcchlhfhvhjhdjhddvjjpmmsrsqrrngrrmvmsvsllrmlrlprpggqzqmmvlvvwrwnrwwzztrzrbrdbrdbbhlhchjchhtthpthphplhhlphpjjsddtppvbpbbmnmgnmmqbbrhhfrfpfjpjgjljcctwtmwtwvvvmsvmmnjmnjmnmqmvvggtzzctzttszsvvjwjqjmjbjvbbshbhghlghhpvpnvvqmmgjmmggqvqtvqtvqtvvmlljbbhdhshnnwqwbbrnnwswmwfmfttjztjzjsjzjdjbdjjzfzdzgzhghgpgqqdzdhhnwnjnnhrnrqqjsqsllbzzzcnzzmjmvmrvrgrfgrrqmrmpmpzmzqqfwqfwfjjhphgpptzzwmmfjjvbjvbbqsbsbcbppjzpjjzpjpvjjzdjjgcgppvrvjrjpjspjppntpnnjlnlrnrdndjnnsbnnppvspprbrddjrdrmrfmmtmtbtntftjjsscvcbvvcnccnbnrbbmlmddhpdptddmrmmjddnjjtztctqqhhttqctqtbbnfnwnmndnzzljlgltgltljtjllwjjbrjbrbpbgppmzpmpggdnnmznnhhmfhhssgtgvvlsljltlppdqqbtbvbsvsswbsbhsbsmmwswbbbbhvbvcbvvsvcsvvmfmlmgmjgmmbqbrbsrrgrsggldlmlttpbpmmtptctqcqvccnhnmnssmvsmmddlclpccfwfdfqfvqfvfqqwmmwrmmbzznpnllfttgcttnbtnnsswttfppnddpdpsphhpttnrtrlrsrmrnmmzhmhnmnwwddmhdhqqhbqqdcclnlfnnzfnzzqfqddczcqcrrjsrrswslltfltffjhhqqfcchqcqpplbltbblbcbncnpccdffcwwqzwqqlwwtjjpbbjvvljvjgvvrsszfsfcssqsvvwsvvdfvdfdpffztzrtzthzhhmwmzwzssjlsldsldsldlvvjdvvnbvnvmvmccvfcvcfcbchhzqqhgqqcbbhdbdpddwcdwwzszsfszzgbgdgzgvgsgnssfqsqjssrppsgppmmpzmppglgqlqqpzpqqfzzzrnzrnrqqzzfwzzffjfzzmnznpznpprhhvcvvvfddcrdrsscnnsznzszbznntwntnrnbnzzrnnlwlggnzzwppmbmrrncnlclzltlblhhlvhvphhsdhsdsqddgzgvvshvvzzbvbrrlplqpptplljcljlvjlvlqvvndnsdddmvddzhzjjgpjpddppzllwtwfwgfgpffmhhzllsttmqqhjhhpvpgpfjsgscnwjmwmtmptwlpfjljwgpgntrlpjfgbjqmcpzgfhrwmznqnsbpptbdrzmdtvvtdqjgrjzlphndhmlchvddglqnqsjqrfqslprsvlqjwqnsmsznptsstpvdntpttslpmqqbsdlqwpjqnzmpblgqmjrvqwsncnzdszgfsghddlnwhwzpgtddgstttvrjfjwwfrgsdjjngljqlqcrzlgsmwngbzvmjwtnqdqcgwmfhsztgrtvvfzbtstmdbqpntdpsszjthqvpbdwswfzvmrcpbgbgdmldfhvdpfsmdzfhwsrpcglsztdwqgbqszcqtqjhgntzvttldqsffftzmllptzhmhpmfsgcchfrnrchnsgcfjbgrqmvrmmhnmlnwtgwhznqfgwnlrqlpjrvfrgzcjwncvlwhpclfzngbgvmrmlzngmqlvvwhbpjzlclgrcnnvnlppqhvlrnpzvmtsbdpfbwgffgzfwvltcfvfdcnfhwvcvclwwbmshhmpgrzgltwjmqczpqzdwfjpqhmwqhvvgnpgtwrjrgwvhthtdrdpnwpbmwstgblwmbfvlwflqmfbcsgwstwvncwfcsmrpcfrrvmlbqhdtdswswfnzhgzlngwsrtlzfcgdppmjnghfrgbdqhqmslhcqddjvsslsjwqqznttqjzdlghnsvqqtwrpfbzjgwnrhhvlnbqmnvcpblzgbzltnrhzpdwvbqbtmctbzgsdjfzswrbqbzgvwjlwtmgcllnmnwcljbhbplpvtgpgjftfrbgpgmhghnjcgjfqmsbqhbgtzbfzgwmfdsgfgmgzsbgdrszfhjttbvcqjzjgbqfgswlmrrhnmnfrptvjtlnvplgznsljzfzmhghlsccnqzflfnmbhshfprhclmtfptcmtnhrjgnngnqnczvcgzzlntftjsbgpgwzbnrhzzfqmznqnzfrvjzsmtpjbswzjlbgpfftzrzbfgdblpwscbqjfrfmfnhhlhjprtlzzvwnwzsnqhnmgwsdprnrblgbclzhthftqzdljspwdzwmwhfmdzmlvqsngppdfsjdprbrhffcvcvzztjjqcffwbrvpzvzfzhjvsfvsnrmjvqmjtrjbmbqsdtjgvtbbzzfmnmrrflgcdtljpmpvqvdbzgbmhjgccgdtplllctzqpfqnsztbwdbmqgfzrcddtmwrgmwsghcfpgqssdjrtqhtjfbpjvjdnzgvpzrhbrhrhcmpbglbbrvdltdpsrwjbjzftccwgnqmnlqlpjwrfdvmlgvgqznlvsmpzsmgjstvqbqpprzlsdndpfbmqcrfgvcfvlfhmpfnnqlcqlnbbgcrrrhbtzwnwmfrnrzvgmqlmqnmnzbwflwzcmncphjqztlrzvpztqhptmfsrppvvlzcfnlrwptgccsjjjscjcwnzssmbcvtzhnscgsbrchbqbrtdzllfvmqfwznfzpzmbfwcdsfhdlddnfbdgqbqjqzdtppshwcvvcjqstdgtgbhmlqlrfrhbvfsszsmbldmwfnfgnjptdslnzwcjgmvbnqcfzjmrslrlllplbhpjcnvzmvrfzwqhmbnrvpqnvcfncgfwqlvcwpwwljssfmswctcmhgtphvjdpfnnzznfbdjcsndrczlgdjhrnrltsgbtqmqcbwlthwcgsbvqbnntcznczpmlmblwdrlmzqdztwsgjthjwtfcpgwbczmdmhttzcwvdzhdfldmwnbnfdcjgvjhrfltjnjhqhzmzrlbncwdnlgshmqhpgsdwbvmvjsfgvgqzqjqdzqzmfmrncfdgrqfrnstvpqwtltnhgrmhgmmnwvlnsfmmbjrdmrnlgfgpqncdpqgvjltpghbgffdppdcfhdqhtvrdnfvcttlrqppfnqtmzpfgsnmjqfrgtbbdzzrccsrzgfjndlrnzqmjjtgldglcpzcrhwfplvdndjtzbpfbbbpfljqnlvrdzbrvczzwdrvzlmslbjsqgqdrltmhwcdpldqldlpctcbjmsvdwlfspzlpgrhdrvjtprcwrmszvzwmmjgvsbfcztmrdgrgshgtggpzszgqhwmhdzjmzsgqsfmqvcgqqwgprgvvqlbfhpwsfbjdqtcfnfhsgzthzhbpwggbnscqbnvcdprsbgllsrdcclqggfgfdrpqlqljfwpmzdhthpczcsqrrm",
			result:   3452,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.input[:10], func(t *testing.T) {
			require.Equal(t, tc.result, tc.f(tc.input, tc.capacity))
		})
	}
}
