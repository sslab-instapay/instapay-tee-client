#ifndef _APP_H_
#define _APP_H_

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

#include "sgx_error.h"       /* sgx_status_t */
#include "sgx_eid.h"     /* sgx_enclave_id_t */

#ifndef TRUE
# define TRUE 1
#endif

#ifndef FALSE
# define FALSE 0
#endif

# define TOKEN_FILENAME   "enclave.token"
# define ENCLAVE_FILENAME "enclave.signed.so"

extern sgx_enclave_id_t global_eid;    /* global enclave id */

#if defined(__cplusplus)
extern "C" {
#endif

int initialize_enclave(void);

/* command.cpp */
void ecall_preset_account_w(unsigned char *addr, unsigned char *seckey);
unsigned char* ecall_create_account_w(void);
unsigned char* ecall_create_channel_w(unsigned int nonce, unsigned char *owner, unsigned char *receiver, unsigned int deposit, unsigned int *sig_len);
int ecall_get_balance_w(unsigned int channel_id);
// void ecall_close_channel_w(void);
// void ecall_eject_w(void);

/* network.cpp */
void ecall_go_pre_update_w(unsigned int payment_num, unsigned int *channel_id, int *amount, unsigned int size);
void ecall_go_post_update_w(unsigned int payment_num, unsigned int *channel_id, int *amount, unsigned int size);
void ecall_go_idle_w(unsigned int payment_num);
void ecall_register_comminfo_w(unsigned int channel_id, unsigned char *ip, unsigned int port);

/* event.cpp */
void ecall_receive_create_channel_w(unsigned int channel_id, unsigned char *owner, unsigned char *receiver, unsigned int deposit);

/* load.cpp */
void ecall_load_account_data_w(unsigned char *addr, unsigned char *seckey);
void ecall_load_channel_data_w(
    unsigned int channel_id,
    unsigned int type,
    unsigned int channel_status,
    unsigned char *my_addr,
    unsigned int my_deposit,
    unsigned int other_deposit,
    unsigned int balance,
    unsigned int locked_balance,
    unsigned char *other_addr,
    unsigned char *other_ip,
    unsigned int other_port);
void ecall_load_payment_data_w(unsigned int payment_num, unsigned int channel_id, int amount);

#if defined(__cplusplus)
}
#endif

#endif /* !_APP_H_ */
